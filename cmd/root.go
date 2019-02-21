package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
)

var dogrcFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dogleash",
	Short: "DataDog CLI tool written in golang",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&dogrcFile, "config", filepath.Join(os.Getenv("HOME"), ".dogrc"), ".dogrc file path")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Check if a file exists
	_, err := os.Stat(dogrcFile)
	if err == nil {
		// read file from dogrcFile
		dogrc, err := ini.Load(dogrcFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: fail to read .dogrc file\n%v", err)
			os.Exit(1)
		}

		viper.SetDefault("api_key", dogrc.Section("Connection").Key("apikey").String())
		viper.SetDefault("app_key", dogrc.Section("Connection").Key("appkey").String())
	}

	// read in environment variables that match
	viper.SetEnvPrefix("DATADOG")
	viper.AutomaticEnv()
}
