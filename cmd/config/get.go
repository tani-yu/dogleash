package config

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "current config name",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile(dogleashFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(viper.GetString("current"))
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
}
