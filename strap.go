package main

import (
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git"
	flag "github.com/spf13/pflag"
)

func main() {
	repo := flag.StringP("repo", "r", "", "Repository to strap")
	outputDir := flag.StringP("output", "o", "", "Output directory for new strap")
	flag.Parse()

	if *repo == "" {
		log.Fatalln("bruh")
	} else if len(strings.Split(*repo, "/")) != 2 {
		log.Fatalln("Malformed input")
	}

	if *outputDir == "" {
		*outputDir = strings.Split(*repo, "/")[1]
	}

	_, err := git.PlainClone(*outputDir, false, &git.CloneOptions{
		URL:      "https://github.com/" + *repo,
		Progress: os.Stdout,
	})

	if err != nil {
		panic(err)
	}
}
