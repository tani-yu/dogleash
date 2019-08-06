//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package dashboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	dd "github.com/tani-yu/dogleash/datadog"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

// dashboardShowAllCmd represents the dashboardShowAllCmd command
var dashboardShowAllCmd = &cobra.Command{
	Use:   "show_all",
	Short: "Show all dashboards data in JSON format",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		printDashboard(cli, "json")
	},
}

func printDashboard(cli *datadog.Client, format string) {
	boards := GetAllDashboards()

	switch format {
	case "json":
		printDashboardAsJSON(cli, boards)
	default:
		log.Fatalf("Error invalid print format: got=%s", format)
	}
}

func printDashboardAsJSON(cli *datadog.Client, boards []datadog.Board) {
	var out bytes.Buffer
	for i, board := range boards {
		dash, err := cli.GetBoard(*board.Id)
		if err != nil {
			log.Fatalf("Error retrieving single dashboard: %s", err)
		}

		jsc, err := json.MarshalIndent(dash, "", "  ")
		if err != nil {
			log.Fatalf("Error unmarshaling responded JSON object: %s\n", err)
		}

		out.Write(jsc)
		if i != len(boards)-1 {
			out.WriteString("\n")
		}
	}
	fmt.Println(out.String())
}

func init() {
	dashboardCmd.AddCommand(dashboardShowAllCmd)
}
