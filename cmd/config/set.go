package config

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a Config for API/APP keys",
	Run: func(cmd *cobra.Command, args []string) {
		// read config
		viper.SetConfigFile(dogleashFile)
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}

		var dogrcExist bool
		_, err = os.Stat(dogrcFile)
		if err != nil {
			dogrcExist = false
		} else {
			dogrcExist = true
		}

		od := viper.Get("organizations")
		if len(args) == 1 {
			// check configs
			/*
				* config exist: update
				* dogrc exist: update
					* dogrc not exist: error
				* others: choose a config
			*/
			for i := 0; i < len(od.([]interface{})); i++ {
				if args[0] == od.([]interface{})[i].(map[interface{}]interface{})["name"] {
					// config exist: update
					UpdateCurrent(args[0])
					return
				} else if args[0] == "dogrc" {
					if dogrcExist {
						// dogrc exist: update
						UpdateCurrent(args[0])
						return
					} else {
						// dogrc not exist: error
						log.Fatalf("\ndoes not exist %s\n", dogrcFile)
					}
				}

			}
			// others: choose a config
			ChooseConfig()
		} else {
			// choose a config
			ChooseConfig()
		}
	},
}

func UpdateCurrent(arg string) {
	viper.Set("current", arg)
	err := viper.WriteConfigAs(dogleashFile)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func ChooseConfig() {
	// list configs
	od := viper.Get("organizations")
	lenod := len(od.([]interface{}))

	fmt.Println("Choose a Config name")
	for i := 0; i < lenod; i++ {
		fmt.Printf("[%d] %s\n", i, od.([]interface{})[i].(map[interface{}]interface{})["name"])
	}
	if _, err := os.Stat(dogrcFile); err == nil {
		fmt.Printf("[%d] dogrc\n", lenod)
	}

	// wait user input
	io.WriteString(os.Stdout, "input number: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatalf("\nChoose valid number.\n%s\n", err)
	}

	// update current
	if n < lenod {
		UpdateCurrent(od.([]interface{})[n].(map[interface{}]interface{})["name"].(string))
	} else if n == lenod {
		UpdateCurrent("dogrc")
	} else {
		log.Fatal("\nChoose valid number.\n")
	}
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
