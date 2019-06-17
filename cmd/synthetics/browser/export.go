package browser

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	dd "github.com/tani-yu/dogleash/datadog"

	"github.com/spf13/cobra"
)

var outputDir string

// browserExportCmd represents the browserExportCmd command
var browserExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export monitors data in JSON format to the specified path",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("Failed to connect Datadog API server: %s\n", err)
		}

		monit, err := cli.GetSyntheticsTests()
		if err != nil {
			log.Fatalf("Error getting all monitors: %s\n", err)
		}

		jsc, err := json.MarshalIndent(monit, "", "  ")
		if err != nil {
			log.Fatalf("Error unmarshaling responded JSON object: %s\n", err)
		}

		baseDir := filepath.Join(outputDir, "monitor")
		if err := os.Mkdir(baseDir, 0755); err != nil {
			log.Fatalf("Error creating monitor datastore directory: %s\n", err)
		}

		file, err := os.Create(filepath.Join(baseDir, "monitor.json"))
		if err != nil {
			log.Fatalf("Error creating json file for all monitors: %s\n", err)
		}
		defer file.Close()
		file.Write(jsc)
	},
}

func init() {
	browserCmd.AddCommand(browserExportCmd)

	browserExportCmd.Flags().StringVarP(&outputDir, "--output-dir", "d", "",
		"already existing destination directory (default is current directory)")
}
