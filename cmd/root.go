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
		Short: "creates .strap.json in current directory",
		Run: func(cmd *cobra.Command, args []string) {
			initProject()
		},
	}

	var config = &cobra.Command{
		Use:   "config",
		Short: "creates ~/.strap.global.json",
		Run: func(cmd *cobra.Command, args []string) {
			// initializer()
		},
	}

	var validate = &cobra.Command{
		Use:   "validate",
		Short: "validates ./.strap.json",
		Run: func(cmd *cobra.Command, args []string) {
			parseProjectCfg()
		},
	}

	var update = &cobra.Command{
		Use:   "update",
		Short: "update current package to x.y+1",
		Run: func(cmd *cobra.Command, args []string) {
			updateProject(cmd)
		},
	}

	update.Flags().StringP("version", "v", "", "version number to update to")

	rootCmd.AddCommand(
		initialize,
		config,
		validate,
		update,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
