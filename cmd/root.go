package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
)

var dogrcFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-datadog",
	Short: "datadogのコンフィグをインポートしたり、エクスポートしたりするためのツール",
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

// FIXME: help サブコマンドの実行時は config 存在確認でコケたりしないよう config の読み込みをスキップしたい
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&dogrcFile, "dogrc", "", ".dogrc file path (default is $HOME/.dogrc)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// reading config in default config path
	path := os.Getenv("HOME") + "/.dogrc"
	dogrc, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read dogrc file: %v", err)
		os.Exit(1)
	}

	viper.SetDefault("api_key", dogrc.Section("Connection").Key("apikey").String())
	viper.SetDefault("app_key", dogrc.Section("Connection").Key("appkey").String())

	// read in environment variables that match
	viper.SetEnvPrefix("DATADOG")
	viper.AutomaticEnv()
}
