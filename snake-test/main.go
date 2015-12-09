package main

import (
	"fmt"
	"os"

	"github.com/breml/mygoplayground/snake-test/cmd"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("snake")
	viper.BindEnv("target")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/snake/")
	viper.AddConfigPath("$HOME/.snake")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
