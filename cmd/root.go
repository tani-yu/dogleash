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
	// check file exist
	var dogrcExist bool
	_, err := os.Stat(dogrcFile)
	if err != nil {
		dogrcExist = false
	} else {
		dogrcExist = true
	}
	var dogleashExist bool
	_, err = os.Stat(dogleashFile)
	if err != nil {
		dogleashExist = false
	} else {
		dogleashExist = true
	}

	// set config api/app keys
	if dogleashExist {
		viper.SetConfigFile(dogleashFile)
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}

		cs := viper.Get("current")
		od := viper.Get("organizations")
		if cs == "dogrc" {
			if dogrcExist {
				dogrc, err := ini.Load(dogrcFile)
				if err != nil {
					log.Fatal(err)
				}
				viper.SetDefault("api_key", dogrc.Section("Connection").Key("apikey").String())
				viper.SetDefault("app_key", dogrc.Section("Connection").Key("appkey").String())
			} else {
				log.Fatal("\ncurrent config is dogrc. but does not exist ~/.dogrc")
			}
		} else {
			for i := 0; i < len(od.([]interface{})); i++ {
				if od.([]interface{})[i].(map[interface{}]interface{})["name"] == cs {
					viper.SetDefault("api_key", od.([]interface{})[i].(map[interface{}]interface{})["keys"].(map[interface{}]interface{})["apikey"])
					viper.SetDefault("app_key", od.([]interface{})[i].(map[interface{}]interface{})["keys"].(map[interface{}]interface{})["appkey"])
				}
			}
		}
	} else {
		log.Fatalf("\ndoes not exist dogleash configfile. [%s]\n", dogleashFile)
	}

	// read in environment variables that match
	viper.SetEnvPrefix("DATADOG")
	viper.AutomaticEnv()
}
