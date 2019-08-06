//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package downtime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

// downtimeScheduleCmd represents the downtimeScheduleCmd command
var downtimeScheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Create downtimes on Datadog by importing JSON object",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		var dts []datadog.Downtime
		for _, inputPath := range args {
			var decoded datadog.Downtime
			raw, err := ioutil.ReadFile(inputPath)
			if err != nil {
				log.Fatalf("Failed to read JSON file: %s\n", err)
			}

			if err := json.Unmarshal(raw, &decoded); err != nil {
				log.Fatalf("JSON Unmarshal error: %s\n", err)
			}
			dts = append(dts, decoded)
		}

		downtimes, err := cli.GetDowntimes()
		if err != nil {
			log.Fatalf("Error getting all downtimes: %s\n", err)
		}

		for _, dt := range dts {
			if checkID(dt, downtimes) {
				out, err := cli.CreateDowntime(&dt)
				if err != nil {
					log.Fatalf("failed to create downtime items: %s\n", err)
				}
				fmt.Printf("CREATE ID: %d\n", *out.Id)
			}
		}
	},
}

func init() {
	downtimeCmd.AddCommand(downtimeScheduleCmd)
}

// Check if there is the same id
func checkID(dt datadog.Downtime, downtimes []datadog.Downtime) bool {
	if dt.Id != nil {
		for _, downtime := range downtimes {
			if *downtime.Id == *dt.Id {
				return false
			}
		}
	}
	return true
}
