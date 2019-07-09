//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

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
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		var synthetics []datadog.SyntheticsTest
		for _, inputPath := range args {
			var decoded []datadog.SyntheticsTest
			raw, err := ioutil.ReadFile(inputPath)
			if err != nil {
				log.Fatalf("Failed to read JSON file: %s\n", err)
			}

			if err := json.Unmarshal(raw, &decoded); err != nil {
				log.Fatalf("JSON Unmarshal error: %s\n", err)
			}
			synthetics = append(synthetics, decoded...)
		}

		syns, err := cli.GetSyntheticsTests()
		if err != nil {
			log.Fatalf("Failed to get synthetics items: %s\n", err)
		}

		for _, synthetic := range synthetics {
			if checkNameAndID(synthetic, syns) {
				maskUnsupportedProperties(&synthetic)
				fmt.Printf("CREATE NAME:%s\n", *synthetic.Name)
				_, err := cli.CreateSyntheticsTest(&synthetic)
				if err != nil {
					log.Fatalf("Failed to create synthetics items: %s\n", err)
				}
			}
		}
	},
}

func init() {
	syntheticsCmd.AddCommand(syntheticsImportCmd)
}

// Check if there is the same id and name
func checkNameAndID(synthetic datadog.SyntheticsTest, syns []datadog.SyntheticsTest) bool {
	for _, syn := range syns {
		if *syn.Name == *synthetic.Name || *syn.MonitorId == *synthetic.MonitorId {
			return false
		}
	}
	return true
}

// maskUnsupportedProperties sets nil in unsuported properites.
// properties below are not allowed by Datadog API:
//     PublicId
//     MonitorId
//     CreatedAt
//     ModifiedAt
func maskUnsupportedProperties(synthetic *datadog.SyntheticsTest) {
	synthetic.PublicId = nil
	synthetic.MonitorId = nil
	synthetic.CreatedAt = nil
	synthetic.ModifiedAt = nil
}
