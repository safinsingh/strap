package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "strap",
	Short: "strap: bootstrap your project templates with ease",
}

// Execute is the command root
func Execute() {
	var initialize = &cobra.Command{
		Use:   "init",
		Short: "Creates .strap.json in current dir",
		Run: func(cmd *cobra.Command, args []string) {
			initProject()
		},
	}

	var config = &cobra.Command{
		Use:   "config",
		Short: "Creates ~/.strap.global.json",
		Run: func(cmd *cobra.Command, args []string) {
			// initializer()
		},
	}

	var validate = &cobra.Command{
		Use:   "validate",
		Short: "Validates ./.strap.json",
		Run: func(cmd *cobra.Command, args []string) {
			parseProjectCfg()
		},
	}

	rootCmd.AddCommand(initialize, config, validate)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}