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

package synthetics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

// syntheticsImportCmd represents the syntheticsImportCmd command
var syntheticsImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Create synthetics on Datadog by importing JSON object",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("failed to connect Datadog API server: %s\n", err)
		}

		var synthetics []datadog.SyntheticsTest
		for _, inputPath := range args {
			var decoded []datadog.SyntheticsTest
			raw, err := ioutil.ReadFile(inputPath)
			if err != nil {
				log.Fatalf("failed to read JSON file: %s\n", err)
			}

			if err := json.Unmarshal(raw, &decoded); err != nil {
				log.Fatalf("JSON Unmarshal error: %s\n", err)
			}
			synthetics = append(synthetics, decoded...)
		}

		syns, err := cli.GetSyntheticsTests()
		if err != nil {
			log.Fatalf("failed to get synthetics items: %s\n", err)
		}

		for _, synthetic := range synthetics {
			if checkNameAndID(synthetic, syns) {
				disallowUnexpectedProperties(&synthetic)
				fmt.Printf("CREATE NAME:%s\n", *synthetic.Name)
				_, err := cli.CreateSyntheticsTest(&synthetic)
				if err != nil {
					log.Fatalf("failed to create synthetics items: %s\n", err)
				}
			}
		}
	},
}

func init() {
	SyntheticsCmd.AddCommand(syntheticsImportCmd)
}

// Check if there is the same id and name
func checkNameAndID(synthetic datadog.SyntheticsTest, syns []datadog.SyntheticsTest) bool {
	for _, syn := range syns {
		if *syn.Name == *synthetic.Name {
			return false
		}
	}
	return true
}

// disallowUnexpectedProperties sets nil in properites.
// properties below are not allowed by Datadog API:
//     PublicId
//     MonitorId
//     CreatedAt
//     ModifiedAt
func disallowUnexpectedProperties(synthetic *datadog.SyntheticsTest) {
	synthetic.PublicId = nil
	synthetic.MonitorId = nil
	synthetic.CreatedAt = nil
	synthetic.ModifiedAt = nil
}
