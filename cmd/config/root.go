package config

import (
	"github.com/tani-yu/dogleash/cmd"
	"github.com/spf13/cobra"
)

var dogrcFile string

var configCmd = &cobra.Command{
	Use: "config",
	Short: "Perform operations related to config of Datadog API/APP keys",
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
}
