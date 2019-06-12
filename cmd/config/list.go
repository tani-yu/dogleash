package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "show list of all config names",
	Run: func(cmd *cobra.Command, args []string) {
		// read config
		viper.SetConfigFile(dogleashFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		err = viper.Unmarshal(&DC)
		if err != nil {
			log.Fatal(err)
		}

		// print configs
		for _, o := range DC.Organizations {
			fmt.Println(o.Name)
		}
		if _, err = os.Stat(dogrcFile); err == nil {
			fmt.Println("dogrc")
		}
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
}
