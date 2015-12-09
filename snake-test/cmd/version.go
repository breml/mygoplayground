package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("snake-test v0.1 -- HEAD")
		} else {
			fmt.Println("v0.1")
		}
	},
}
