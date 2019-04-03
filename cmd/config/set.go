package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use: "set",
	Short: "Set Config Organization for API/APP keys",
	Run: func(cmd *cobra.Command, args []string) {
		var file *os.File
		if len(args) == 1 {
			dogrcFile = filepath.Join(os.Getenv("HOME"), ".dogrc.d/current")
			_, err := os.Stat(dogrcFile)
			if err != nil {
				file, err = os.OpenFile(dogrcFile, os.O_WRONLY|os.O_CREATE, 0644)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
					os.Exit(1)
				}
				defer file.Close()
			} else {
				file, err = os.OpenFile(dogrcFile, os.O_WRONLY|os.O_TRUNC, 0644)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
					os.Exit(1)
				}
				defer file.Close()
			}
			fmt.Fprintf(file, args[0])
		} else {
			fmt.Fprintf(os.Stderr, "set <config>\n")
			os.Exit(1)
		}
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
