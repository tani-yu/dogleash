package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tani-yu/dogleash/cmd"
)

// dogleash config current and orgs
type DogleashConfig struct {
	Organizations []Organization
	Current       string `mapstructure:"current"`
}

// orgs
type Organization struct {
	Name   string `mapstructure:"name"`
	APIKey string `mapstructure:"apikey"`
	APPKey string `mapstructure:"appkey"`
}

// DC dogleash config
var DC DogleashConfig

// config path
var dogrcFile = filepath.Join(os.Getenv("HOME"), ".dogrc")
var dogleashFile = filepath.Join(os.Getenv("HOME"), ".config/dogleash/config.yml")

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Perform operations related to config of Datadog API/APP keys",
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
}
