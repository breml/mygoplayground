package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose bool
	target  string
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "target")
	viper.BindPFlag("target", RootCmd.PersistentFlags().Lookup("target"))
	viper.SetDefault("target", "apache")
}

var RootCmd = &cobra.Command{
	Use:   "snake-test",
	Short: "snake-test is a test command",
	Long:  `snake-test command long description`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("snake-test:", viper.GetString("target"))
		}
	},
}
