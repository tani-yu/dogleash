package config

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var configListCmd = &cobra.Command{
        Use: "list",
        Short: "list Organization configs",
        Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(filepath.Join(os.Getenv("HOME"), ".dogrc.d"))
		if err != nil {
		        fmt.Fprintf(os.Stderr, "%v\n", err)
		        os.Exit(1)
		}
		for _, f := range files {
			if f.Name() == "current" {
				continue
			}
			fmt.Println(f.Name())
		}
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
}
