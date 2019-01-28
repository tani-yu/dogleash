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
	"os"

	"github.com/tani-yu/dogleash/cmd/auth"
	"github.com/spf13/cobra"
)

var outputPath string

// compute_checkCmd represents the compute_check command
var monitorShowAllCmd = &cobra.Command{
	Use:   "show_all",
	Short: "すべてのモニターをjson形式で出力します",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := auth.GetDDClient()
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
		}

		monit, err := cli.GetMonitors()
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
		}

		jsc, _ := json.MarshalIndent(monit, "", "  ")
		if outputPath == "" {
			fmt.Println(string(jsc))
		} else {
			os.Mkdir(outputPath+"monitor/", 0755)
			file, err := os.Create(outputPath + "monitor/monitor.json")
			if err != nil {
				log.Fatalf("fatal: %s\n", err)
			}
			defer file.Close()
			file.Write(jsc)
		}
	},
}

// VM の状態と並行ジョブの実行状態を更新する
//
// goroutine として呼び出される
func checkJobStatus() {
	return
}

func init() {
	monitorCmd.AddCommand(monitorShowAllCmd)
	monitorShowAllCmd.Flags().StringVarP(&outputPath, "output", "p", "",
		"指定された場所にJSONファイルを保存")
}
