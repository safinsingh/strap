package main

import (
	"os"

	"github.com/go-git/go-git"
)

func cloneRepo(repo, outputDir string) {
	if _, err := git.PlainClone(outputDir, false, &git.CloneOptions{
		URL:      "https://github.com/" + repo,
		Progress: os.Stdout,
	}); err != nil {
		errorPrint(err, "Error cloning!")
	}

	if err := os.RemoveAll(outputDir + "/.git"); err != nil {
		errorPrint(err, "Error recursively removing .git!")
	}
}
