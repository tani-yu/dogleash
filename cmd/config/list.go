package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "list configs",
	Run: func(cmd *cobra.Command, args []string) {
		// read config
		viper.SetConfigFile(dogleashFile)
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
		// print configs
		od := viper.Get("organizations")
		for i := 0; i < len(od.([]interface{})); i++ {
			fmt.Printf("%s\t", od.([]interface{})[i].(map[interface{}]interface{})["name"])
		}
		if _, err = os.Stat(dogrcFile); err != nil {
			fmt.Println()
		} else {
			fmt.Println("dogrc")
		}
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
}
