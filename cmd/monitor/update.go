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
	"fmt"
	"io/ioutil"
	"log"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

// monitorUpdateCmd represents the monitorUpdateCmd command
var monitorUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update monitors on Datadog by importing JSON object",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("failed to connect Datadog API server: %s\n", err)
		}

		var monits []datadog.Monitor
		for _, inputPath := range args {
			var decoded []datadog.Monitor
			raw, err := ioutil.ReadFile(inputPath)
			if err != nil {
				log.Fatalf("failed to read JSON file: %s\n", err)
			}

			if err := json.Unmarshal(raw, &decoded); err != nil {
				log.Fatalf("JSON Unmarshal error: %s\n", err)
			}
			monits = append(monits, decoded...)
		}

		mons, err := cli.GetMonitors()
		if err != nil {
			log.Fatalf("failed to get monitoring items: %s\n", err)
		}

		for _, monit := range monits {
			if checkID(monit, mons) {
				fmt.Printf("UPDATE  ID:%d, NAME:%s\n", *monit.Id, *monit.Name)
				err := cli.UpdateMonitor(&monit)
				if err != nil {
					log.Fatalf("failed to update monitoring items: %s\n", err)
				}
			}
		}
	},
}

func init() {
	monitorCmd.AddCommand(monitorUpdateCmd)
}

// Check if there is the same id
func checkID(monit datadog.Monitor, mons []datadog.Monitor) bool {
	for _, mon := range mons {
		if *mon.Id == *monit.Id {
			return true
		}
	}
	return false
}
