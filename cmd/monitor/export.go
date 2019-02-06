// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitor

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
)

var targetDir string

// monitorExportCmd represents the monitorExportCmd command
var monitorExportCmd = &cobra.Command{
	Use:   "export",
	Short: "export monitor data in JSON format to the specified path",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
		}

		monit, err := cli.GetMonitors()
		if err != nil {
			log.Fatalf("Error getting all monitors: %s\n", err)
		}

		jsc, err := json.MarshalIndent(monit, "", "  ")
		if err != nil {
			log.Fatalf("Error unmarshaling responded JSON object: %s\n", err)
		}

		baseDir := filepath.Join(targetDir, "monitor")
		if err := os.Mkdir(baseDir, 0755); err != nil {
			log.Fatalf("Error creating monitor datastore directory: %s\n", err)
		}

		file, err := os.Create(filepath.Join(baseDir, "monitor.json"))
		if err != nil {
			log.Fatalf("Error creating json file for all monitors: %s\n", err)
		}
		defer file.Close()
		file.Write(jsc)
	},
}

func init() {
	monitorCmd.AddCommand(monitorExportCmd)

	monitorExportCmd.Flags().StringVarP(&targetDir, "--target-dir", "d", "",
		"already existing destination directory (default is current directory)")
}
