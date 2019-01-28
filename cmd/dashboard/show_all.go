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

package dashboard

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tani-yu/dogleash/cmd/auth"
	"github.com/spf13/cobra"
)

var outputPath string

// compute_checkCmd represents the compute_check command
var dashboardShowAllCmd = &cobra.Command{
	Use:   "show_all",
	Short: "すべてのdashboardをjson形式で出力します",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := auth.GetDDClient()
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
		}

		boards, err := cli.GetDashboards()
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
		}

		for _, board := range boards {
			dash, _ := cli.GetDashboard(*board.Id)
			jsc, _ := json.MarshalIndent(dash, "", "  ")
			if outputPath == "" {
				fmt.Println(string(jsc))
			} else {
				os.Mkdir(outputPath+"dashboard/", 0755)
				file, err := os.Create(outputPath + "dashboard/" + validationFileName(board.Title) + ".json")
				if err != nil {
					log.Fatalf("fatal: %s\n", err)
				}
				defer file.Close()
				file.Write(jsc)
			}
		}
	},
}

// linuxで使えない文字列を変換（graphのtitleの中に元のデータは残るはず）
func validationFileName(s *string) string {
	res := strings.Replace(*s, " ", "_", -1)
	return strings.Replace(res, "/", "", -1)
}

func init() {
	dashboardCmd.AddCommand(dashboardShowAllCmd)

	dashboardShowAllCmd.Flags().StringVarP(&outputPath, "output", "p", "",
		"指定された場所にJSONファイルを保存")
}
