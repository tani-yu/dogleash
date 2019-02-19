package monitor

import (
	"encoding/json"
	"fmt"
	"log"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
)

// monitorShowAllCmd represents the monitorShowAllCmd command
var monitorShowAllCmd = &cobra.Command{
	Use:   "show_all",
	Short: "show all monitor info in JSON format",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		monit, err := cli.GetMonitors()
		if err != nil {
			log.Fatalf("Error getting all monitors: %s\n", err)
		}

		jsc, err := json.MarshalIndent(monit, "", "  ")
		if err != nil {
			log.Fatalf("Error unmarshaling responded JSON object: %s\n", err)
		}

		fmt.Println(string(jsc))
	},
}

func init() {
	monitorCmd.AddCommand(monitorShowAllCmd)
}
