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
		Short: "Initialize a local/global configuration",
		Run: func(cmd *cobra.Command, args []string) {
			initSwitch(cmd)
		},
	}

	var validate = &cobra.Command{
		Use:   "validate",
		Short: "Validates local/global configuration",
		Run: func(cmd *cobra.Command, args []string) {
			parseCfgSwitch(cmd)
		},
	}

	var update = &cobra.Command{
		Use:   "update",
		Short: "update local package to x.y+1",
		Run: func(cmd *cobra.Command, args []string) {
			updateProject(cmd)
		},
	}

	var get = &cobra.Command{
		Use:   "get",
		Short: "Get a remote repository",
		Run: func(cmd *cobra.Command, args []string) {
			getRepo(cmd)
		},
	}

	get.Flags().StringP("repo", "r", "", "remote repository to clone")
	get.Flags().StringP("output", "o", "", "output directory for clone")

	validate.Flags().BoolP("global", "g", false, "affect global or local settings")
	initialize.Flags().BoolP("global", "g", false, "affect global or local settings")
	update.Flags().StringP("version", "v", "", "version number to update to")

	rootCmd.AddCommand(
		initialize,
		validate,
		update,
		get,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
