package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

func parseFlags() (string, string) {
	repo := flag.StringP("repo", "r", "", "Repository to strap")
	outputDir := flag.StringP("output", "o", "", "Output directory for new strap")

	flag.Parse()

	if *repo == "" {
		failPrint("Invalid arguments\n")
		fmt.Println("Usage of " + os.Args[0] + ":")
		flag.PrintDefaults()
		os.Exit(1)

	} else if len(strings.Split(*repo, "/")) != 2 {
		failPrint("Malformed repository name\n")
		fmt.Println("Usage of " + os.Args[0] + ":")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *outputDir == "" {
		*outputDir = strings.Split(*repo, "/")[1]
	}

	return *repo, *outputDir
}
