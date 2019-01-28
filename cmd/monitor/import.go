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
	"os"

	"github.com/tani-yu/dogleash/cmd/auth"
	"github.com/spf13/cobra"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

var inputPath string

// compute_checkCmd represents the compute_check command
var monitorImportCmd = &cobra.Command{
	Use:   "import",
	Short: "json形式のファイルを読み込んでdatadogにdashboardを作成します",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := auth.GetDDClient()
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
		}

		raw, err := ioutil.ReadFile(inputPath)
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
			os.Exit(1)
		}

		var monits []datadog.Monitor
		json.Unmarshal(raw, &monits)

		for _, monit := range monits {
			if checkNameAndID(monit, cli) {
				fmt.Printf("CREATE  ID:%d, NAME:%s\n", *monit.Id, *monit.Name)
				_, err := cli.CreateMonitor(&monit)
				if err != nil {
					log.Fatalf("fatal: %s\n", err)
				}
			}
		}
	},
}

func init() {
	monitorCmd.AddCommand(monitorImportCmd)
	// directory指定できるように
	monitorImportCmd.Flags().StringVarP(&inputPath, "input", "i", "",
		"JSONファイルを指定")
}

// Check if there is the same id and name
func checkNameAndID(monit datadog.Monitor, cli *datadog.Client) bool {
	mons, err := cli.GetMonitors()
	for _, mon := range mons {
		if *mon.Id == *monit.Id || *mon.Name == *monit.Name {
			return false
		}
	}
	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}

	return true
}