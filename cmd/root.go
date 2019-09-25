package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
)

// dogleash config current and orgs
type DogleashConfig struct {
	Organizations []Organization
	Current       string `mapstructure:"current"`
}

// orgs
type Organization struct {
	Name   string `mapstructure:"name"`
	APIKey string `mapstructure:"apikey"`
	APPKey string `mapstructure:"appkey"`
}

// DC dogleash config
var DC DogleashConfig

// config path
var dogrcFile = filepath.Join(os.Getenv("HOME"), ".dogrc")
var dogleashFile = filepath.Join(os.Getenv("HOME"), ".config/dogleash/config.yml")

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
	}

	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "config":
			return
		case "help":
			return
		}
	}

	initConfigDDKey()
}

func initConfigDDKey() {
	_, err := os.Stat(dogleashFile)
	if err != nil {
		log.Fatalf("\nDogleash config file does not exist. [%s]\n", dogleashFile)
	}

	viper.SetConfigFile(dogleashFile)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&DC)
	if err != nil {
		log.Fatal(err)
	}

	if DC.Current == "dogrc" {
		_, err := os.Stat(dogrcFile)
		if err != nil {
			log.Fatal("\nCurrent config is set as \"dogrc\", but ~/.dogrc does not exist.")
		}

		dogrc, err := ini.Load(dogrcFile)
		if err != nil {
			log.Fatal(err)
		}

		viper.SetDefault("api_key", dogrc.Section("Connection").Key("apikey").String())
		viper.SetDefault("app_key", dogrc.Section("Connection").Key("appkey").String())
	} else {
		for _, o := range DC.Organizations {
			if o.Name == DC.Current {
				viper.SetDefault("api_key", o.APIKey)
				viper.SetDefault("app_key", o.APPKey)
			}
		}
	}

	// read in environment variables that match
	viper.SetEnvPrefix("DATADOG")
	viper.AutomaticEnv()
}
