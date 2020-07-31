package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

	infoPrint("Creating ./.strap.json...")

	if err := ioutil.WriteFile("./.strap.json", []byte(jsonTemplate), 0644); err != nil {
		log.Fatalln("Failed to write to ./.strap.json. Run in verbose mode for more details.")
	} else {
		successPrint("Successfully wrote to ./.strap.json!")
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

	newVersion := major + "." + strconv.Itoa(minor+1)

	if len(args) == 0 {
		infoPrint("No version number specified. Bumping to " + newVersion + ".")
		config.BumpVersion(newVersion)

		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatalln(errors.Wrap(err, "Internal error: failed to marshal config struct"))
		}

		if err := ioutil.WriteFile("./.strap.json", data, 644); err != nil {
			log.Fatalln("Failed to write to ./.strap.json. Run in verbose mode for more details.")
		} else {
			successPrint("Successfully bumped version to " + newVersion + "!")
		}
	}
}
