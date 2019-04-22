package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Config name is set",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile(dogleashFile)
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
		fmt.Println(viper.GetString("current"))
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
}
