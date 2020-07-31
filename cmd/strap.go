package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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
		failPrint("Failed to write to ./.strap.json. Run in verbose mode for more details.")
	} else {
		successPrint("Successfully wrote to ./.strap.json!")
	}
}

func parseProjectCfg() {
	data, err := ioutil.ReadFile("./.strap.json")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to read ./.strap.json"))
	}

	var cfg ProjectConfig
	if err := json.Unmarshal([]byte(data), &cfg); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg)
}
