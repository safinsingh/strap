package main

import (
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

func errorPrint(err error, message string) {
	red := color.New(color.FgRed, color.Bold)
	red.Println(errors.Wrap(err, "[-]"+message))
}

func failPrint(message string) {
	red := color.New(color.FgRed, color.Bold)
	red.Println("[-] " + message)
}
