package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Config Organization for API/APP keys",
	Run: func(cmd *cobra.Command, args []string) {
		dogrcFile = filepath.Join(os.Getenv("HOME"), ".dogrc.d/current")
		file, err := os.Open(dogrcFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
}
