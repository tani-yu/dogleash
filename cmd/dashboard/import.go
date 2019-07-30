//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package dashboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	dd "github.com/tani-yu/dogleash/datadog"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

// dashboardImportCmd represents the dashboardImportCmd command
var dashboardImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Create dashbaords on Datadog by importing JSON object",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: After monitor import function created, it will create.
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		var bos []datadog.Board
		for _, inputPath := range args {
			var decoded datadog.Board
			raw, err := ioutil.ReadFile(inputPath)
			if err != nil {
				log.Fatalf("Failed to read JSON file: %s\n", err)
			}

			if err := json.Unmarshal(raw, &decoded); err != nil {
				log.Fatalf("JSON Unmarshal error: %s\n", err)
			}
			bos = append(bos, decoded)
		}

		boards := GetAllDashboards()

		for _, bo := range bos {
			if checkID(bo, boards) {
				fmt.Printf("CREATE  ID:%s, NAME:%s\n", *bo.Id, *bo.Title)
				_, err := cli.CreateBoard(&bo)
				if err != nil {
					log.Fatalf("failed to create dashboard items: %s\n", err)
				}
			} else {
				fmt.Printf("dashboard ID: \"%s\" is already exists\n", *bo.Id)
			}
		}
	},
}

func init() {
	dashboardCmd.AddCommand(dashboardImportCmd)
}

// Check if there is the same id and name
func checkID(bo datadog.Board, boards []datadog.Board) bool {
	for _, board := range boards {
		if *board.Id == *bo.Id {
			return false
		}
	}
	return true
}
