//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package downtime

import (
	"github.com/tani-yu/dogleash/cmd"

	"github.com/spf13/cobra"
)

// downtimeCmd represents the downtimeCmd
var downtimeCmd = &cobra.Command{
	Use:   "downtime",
	Short: "Perform operations related to downtime of Datadog",
}

func init() {
	cmd.RootCmd.AddCommand(downtimeCmd)
}
