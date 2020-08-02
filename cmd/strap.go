package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func initProject() {
	jsonTemplate := `{
	"version": "1.0",
	"commands": {
		"default": {
			"steps": [],
			"entrypoint": ["./exe"]
		}
	}
}`

	if !fileExists("./.strap.json") {
		infoPrint("Creating ./.strap.json...")
		if err := ioutil.WriteFile("./.strap.json", []byte(jsonTemplate), 0644); err != nil {
			log.Fatalln(errors.Wrap(err, "Failed to write to ./.strap.json."))
		} else {
			successPrint("Successfully wrote to ./.strap.json!")
		}
	} else {
		failPrint("./.strap.json alread exists. Please remove it if you would like to re-initialize your project")
	}
}

func parseProjectCfg() ProjectConfig {
	data, err := ioutil.ReadFile("./.strap.json")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to read ./.strap.json"))
	}

	var cfg ProjectConfig
	if err := json.Unmarshal([]byte(data), &cfg); err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to unmarshal JSON into struct. Please check your config file."))
	}

	successPrint("Successfully validated configuration!")
	return cfg
}

func updateProject(cmd *cobra.Command) {
	config := parseProjectCfg()
	currentVersion := config.Version

	major := strings.Split(currentVersion, ".")[0]
	minorStr := strings.Split(currentVersion, ".")[1]

	minor, err := strconv.Atoi(minorStr)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to convert "+strings.Split(currentVersion, ".")[1]+"to integer"))
	}

	context, err := cmd.Flags().GetString("version")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to parse version flag"))
	}

	if context == "" {
		newVersion := major + "." + strconv.Itoa(minor+1)

		infoPrint("No version number specified. Bumping " + config.Name + " to version " + newVersion + ".")
		config.BumpVersion(newVersion)

		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatalln(errors.Wrap(err, "Internal error: failed to marshal config struct"))
		}

		if err := ioutil.WriteFile("./.strap.json", data, 644); err != nil {
			log.Fatalln("Failed to write to ./.strap.json. Run in verbose mode for more details.")
		} else {
			successPrint("Successfully bumped" + config.Name + " to version " + newVersion + "!")
		}
	} else {
		newVersionSlice := strings.Split(context, ".")

		if len(newVersionSlice) == 2 {
			if _, err := strconv.Atoi(newVersionSlice[0]); err == nil {
				if _, err2 := strconv.Atoi(newVersionSlice[1]); err2 == nil {
					successPrint("Valid version number " + context + " has been supplied.")
				} else {
					log.Fatalln(errors.Wrap(err, "Invalid minor version number supplied."))
				}
			} else {
				log.Fatalln(errors.Wrap(err, "Invalid major version number supplied."))
			}
		} else {
			failPrint("Invalid version number provided. Please use the format x.y.")
			os.Exit(1)
		}

		infoPrint("Bumping " + config.Name + " to version " + context + ".")
		config.BumpVersion(context)

		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatalln(errors.Wrap(err, "Internal error: failed to marshal config struct"))
		}

		if err := ioutil.WriteFile("./.strap.json", data, 644); err != nil {
			log.Fatalln(errors.Wrap(err, "Failed to write to ./.strap.json."))
		} else {
			successPrint("Successfully bumped " + config.Name + " to version " + context + "!")
		}
	}
}

func getOutputDir(name, output string) string {
	if output == "" {
		return name
	}
	return output
}

func defaultRun(cmd *cobra.Command) {
	repoFlag, err := cmd.Flags().GetString("repo")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to parse repo flag"))
	}
	repo := strings.Split(repoFlag, "/")[1]

	outputFlag, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to parse output flag"))
	}

	if repoFlag == "" {
		cmd.Help()
		os.Exit(0)
	} else {
		if _, err := git.PlainClone(getOutputDir(repo, outputFlag), false, &git.CloneOptions{
			URL:      "https://github.com/" + repoFlag,
			Progress: os.Stdout,
		}); err != nil {
			log.Fatalln(errors.Wrap(err, "Error cloning repository "+repoFlag))
		}

		if err := os.RemoveAll(getOutputDir(repo, outputFlag) + "/.git"); err != nil {
			log.Fatalln(errors.Wrap(err, "Error recursively removing .git"))
		}
	}
}
