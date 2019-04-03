package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"bufio"

	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var dogrcFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dogleash",
	Short: "Datadog CLI tool written in golang",
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
}

func initConfig() {
	// Check args
	if len(os.Args) == 1 {
		return
	} else if len(os.Args) != 2 {
		switch os.Args[1] {
			case "config":
				return
			case "help":
				return
			default:
				initConfigDDKey()
		}
	}
}

func initConfigDDKey() {
	// Check if a file exists
	dogrcFile = filepath.Join(os.Getenv("HOME"), ".dogrc.d/current")
	_, err := os.Stat(dogrcFile)
	if err != nil {
		fmt.Println("Usage: dogleash config help")
		os.Exit(1)
	} else {
		// read file from dogrcFile
		file, err := os.Open(dogrcFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: fail to read .dogrc.d/current\n%v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			dogrcFile = filepath.Join(os.Getenv("HOME"), ".dogrc.d/", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: fail to read .dogrc.d/%s\n%v\n", scanner.Text(), err)
			os.Exit(1)
		}
		dogrc, err := ini.Load(dogrcFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: fail to read .dogrc file\n%v\n", err)
			os.Exit(1)
		}

		viper.SetDefault("api_key", dogrc.Section("Connection").Key("apikey").String())
		viper.SetDefault("app_key", dogrc.Section("Connection").Key("appkey").String())
	}

	// read in environment variables that match
	viper.SetEnvPrefix("DATADOG")
	viper.AutomaticEnv()
}
