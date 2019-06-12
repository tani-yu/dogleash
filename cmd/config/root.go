package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tani-yu/dogleash/cmd"
)

// DC dogleash config
var DC cmd.DogleashConfig

// config path
var dogrcFile = filepath.Join(os.Getenv("HOME"), ".dogrc")
var dogleashFile = filepath.Join(os.Getenv("HOME"), ".config/dogleash/config.yml")

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage config files of Datadog API/APP keys",
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
}
