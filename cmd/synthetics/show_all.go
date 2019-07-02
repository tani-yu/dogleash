package synthetics

import (
	"encoding/json"
	"fmt"
	"log"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
)

// syntheticsShowAllCmd represents the syntheticsShowAllCmd command
var syntheticsShowAllCmd = &cobra.Command{
	Use:   "show_all",
	Short: "Show all synthetics data in JSON format",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		synthetics, err := cli.GetSyntheticsTests()
		if err != nil {
			log.Fatalf("Error getting all synthetics: %s\n", err)
		}

		jsc, err := json.MarshalIndent(synthetics, "", "  ")
		if err != nil {
			log.Fatalf("Error unmarshaling responded JSON object: %s\n", err)
		}

		fmt.Println(string(jsc))
	},
}

func init() {
	syntheticsCmd.AddCommand(syntheticsShowAllCmd)
}
