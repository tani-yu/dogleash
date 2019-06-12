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

		err = viper.Unmarshal(&DC)
		if err != nil {
			log.Fatal(err)
		}
		if len(args) == 1 {
			// check configs
			/*
				* config exist: update
				* dogrc exist: update
					* dogrc not exist: error
				* others: choose a config
			*/
			for _, o := range DC.Organizations {
				if args[0] == o.Name {
					// config exist: update
					UpdateCurrent(args[0])
					return
				} else if args[0] == "dogrc" {
					if dogrcExist {
						// dogrc exist: update
						UpdateCurrent(args[0])
					} else {
						// dogrc not exist: error
						log.Fatalf("\ndoes not exist %s\n", dogrcFile)
					}
					return
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

// UpdateCurrent set current config
func UpdateCurrent(arg string) {
	viper.Set("current", arg)
	err := viper.WriteConfigAs(dogleashFile)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// ChooseConfig make user to choose config from list
func ChooseConfig() {
	// list configs
	err := viper.Unmarshal(&DC)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Choose a Config name")
	for i, o := range DC.Organizations {
		fmt.Printf("[%d] %s\n", i, o.Name)
	}
	if _, err := os.Stat(dogrcFile); err == nil {
		fmt.Printf("[%d] dogrc\n", len(DC.Organizations))
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
	if n < len(DC.Organizations) {
		for i, o := range DC.Organizations {
			if i == n {
				UpdateCurrent(o.Name)
			}
		}
	} else if n == len(DC.Organizations) {
		UpdateCurrent("dogrc")
	} else {
		log.Fatal("\nChoose valid number.\n")
	}
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
