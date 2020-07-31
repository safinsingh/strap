package cmd

import "github.com/fatih/color"

func successPrint(text string) {
	green := color.New(color.FgGreen, color.Bold)
	green.Println("[+] " + text)
}

func failPrint(text string) {
	red := color.New(color.FgRed, color.Bold)
	red.Println("[-] " + text)
}

func warnPrint(text string) {
	yellow := color.New(color.FgYellow, color.Bold)
	yellow.Println("[!] " + text)
}

func infoPrint(text string) {
	blue := color.New(color.FgBlue, color.Bold)
	blue.Println("[$] " + text)
}
