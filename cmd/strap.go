package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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

func updateProject(args []string) {
	config := parseProjectCfg()
	currentVersion := config.Version

	major := strings.Split(currentVersion, ".")[0]
	minorStr := strings.Split(currentVersion, ".")[1]

	minor, err := strconv.Atoi(minorStr)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to convert "+strings.Split(currentVersion, ".")[1]+"to integer"))
	}

	if len(args) == 0 {
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
	} else if len(args) == 1 {
		newVersionSlice := strings.Split(args[0], ".")

		if len(newVersionSlice) == 2 {
			if _, err := strconv.Atoi(newVersionSlice[0]); err == nil {
				if _, err2 := strconv.Atoi(newVersionSlice[1]); err2 == nil {
					successPrint("Valid version number " + args[0] + " has been supplied.")
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

		infoPrint("Bumping " + config.Name + " to version " + args[0] + ".")
		config.BumpVersion(args[0])

		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatalln(errors.Wrap(err, "Internal error: failed to marshal config struct"))
		}

		if err := ioutil.WriteFile("./.strap.json", data, 644); err != nil {
			log.Fatalln("Failed to write to ./.strap.json. Run in verbose mode for more details.")
		} else {
			successPrint("Successfully bumped " + config.Name + " to version " + args[0] + "!")
		}
	}
}
