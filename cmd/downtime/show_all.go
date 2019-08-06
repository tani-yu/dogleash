//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package downtime

import (
	"encoding/json"
	"fmt"
	"log"

	dd "github.com/tani-yu/dogleash/datadog"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"

	"github.com/spf13/cobra"
)

var activeonly bool

// downtimeShowAllCmd represents the downtimeShowAllCmd command
var downtimeShowAllCmd = &cobra.Command{
	Use:   "show_all",
	Short: "Show all downtime data in JSON format",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		downtimes, err := cli.GetDowntimes()
		if err != nil {
			log.Fatalf("Error getting all downtimes: %s\n", err)
		}

		var dts []datadog.Downtime
		if activeonly {
			for _, downtime := range downtimes {
				if *downtime.Active {
					dts = append(dts, downtime)
				}
			}
		} else {
			dts = downtimes
		}

		jsc, err := json.MarshalIndent(dts, "", "  ")
		if err != nil {
			log.Fatalf("Error unmarshaling responded JSON object: %s\n", err)
		}

		fmt.Println(string(jsc))
	},
}

func init() {
	downtimeCmd.AddCommand(downtimeShowAllCmd)

	downtimeShowAllCmd.Flags().BoolVarP(&activeonly, "active-only", "a", false,
		"show downtimes that are active. Default shows all downtimes that are active or not.")
}
