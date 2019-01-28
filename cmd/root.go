package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var profileName string

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
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.datadog/config.yaml)")
	RootCmd.PersistentFlags().StringVar(&profileName, "profile", "default", "使用するプロファイル名")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.Printf(cfgFile)
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("config")          // name of config file (without extension)
	viper.AddConfigPath("$HOME/.datadog/") // adding home directory as first search path
	viper.AutomaticEnv()                   // read in environment variables that match

	err := viper.ReadInConfig() // 設定ファイルを探索して読み取る
	if err != nil {             // 設定ファイルの読み取りエラー対応
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}
}

func ProfileName() string {
	return profileName
}
