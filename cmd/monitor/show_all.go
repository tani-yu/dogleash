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
	"log"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
)

// monitorShowAllCmd represents the monitorShowAllCmd command
var monitorShowAllCmd = &cobra.Command{
	Use:   "show_all",
	Short: "show all monitor info in JSON format",
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

		fmt.Println(string(jsc))
	},
}

func init() {
	monitorCmd.AddCommand(monitorShowAllCmd)
}
