//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package downtime

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	dd "github.com/tani-yu/dogleash/datadog"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

// downtimeCancelCmd represents the downtimeCancelCmd command
var downtimeCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel registered downtimes on Datadog",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		downtimes, err := cli.GetDowntimes()
		if err != nil {
			log.Fatalf("Error getting all downtimes: %s\n", err)
		}
		for _, id := range args {
			i, err := strconv.Atoi(id)
			if err != nil {
				log.Fatalf("internale error.\n%s\n", err)
			}
			if searchID(i, downtimes) {
				err := cli.DeleteDowntime(i)
				if err != nil {
					log.Fatalf("Failed to cancel a downtime. ID :%d\n%s\n", i, err)
				}
			}
		}
	},
}

func init() {
	downtimeCmd.AddCommand(downtimeCancelCmd)
}

// Check if there is the same id
func searchID(id int, downtimes []datadog.Downtime) bool {
	for _, downtime := range downtimes {
		if *downtime.Id == id {
			return true
		}
	}
	fmt.Printf("ID: %d is not registered on Datadog Downtimes.\n", id)
	return false
}
