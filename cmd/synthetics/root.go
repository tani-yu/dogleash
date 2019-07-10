//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package synthetics

import (
	"github.com/tani-yu/dogleash/cmd"

	"github.com/spf13/cobra"
)

// syntheticCmd represents the syntheticsCmd
var syntheticsCmd = &cobra.Command{
	Use:   "synthetics",
	Short: "Perform operations related to synthetics of Datadog",
}

func init() {
	cmd.RootCmd.AddCommand(syntheticsCmd)
}
