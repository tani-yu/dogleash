//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package synthetics

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
)

var outputDir string

// syntheticsExportCmd represents the syntheticsExportCmd command
var syntheticsExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export synthetics data in JSON format to the specified path",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		syns, err := cli.GetSyntheticsTests()
		if err != nil {
			log.Fatalf("Error getting all synthetics: %s\n", err)
		}

		jsc, err := json.MarshalIndent(syns, "", "  ")
		if err != nil {
			log.Fatalf("Error unmarshaling responded JSON object: %s\n", err)
		}

		baseDir := filepath.Join(outputDir, "synthetics")
		if err := os.Mkdir(baseDir, 0755); err != nil {
			log.Fatalf("Error creating synthetics datastore directory: %s\n", err)
		}

		file, err := os.Create(filepath.Join(baseDir, "synthetics.json"))
		if err != nil {
			log.Fatalf("Error creating json file for all synthetics: %s\n", err)
		}
		defer file.Close()
		file.Write(jsc)
	},
}

func init() {
	syntheticsCmd.AddCommand(syntheticsExportCmd)

	syntheticsExportCmd.Flags().StringVarP(&outputDir, "--output-dir", "d", "",
		"already existing destination directory (default is current directory)")
}
