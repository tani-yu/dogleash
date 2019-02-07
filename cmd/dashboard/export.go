package dashboard

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	dd "github.com/tani-yu/dogleash/datadog"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

var targetDir string

// dashboardExportCmd represents the dashboardExportCmd command
var dashboardExportCmd = &cobra.Command{
	Use:   "export",
	Short: "export dashboard info in JSON format to the specified path",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := dd.NewDDClient()
		if err != nil {
			log.Fatalf("fatal: %s\n", err)
		}

		exportDashboard(cli, targetDir, "json")
	},
}

func exportDashboard(cli *datadog.Client, path, format string) {
	boards, err := cli.GetDashboards()
	if err != nil {
		log.Fatalf("Error retrieving all dashboards: %s\n", err)
	}

	switch format {
	case "json":
		exportDashboardAsJSON(cli, boards, path)
	default:
		log.Fatalf("Error invalid export format: got=%s", format)
	}
}

func exportDashboardAsJSON(cli *datadog.Client, boards []datadog.DashboardLite, path string) {
	baseDir := filepath.Join(path, "dashboard")
	if err := os.Mkdir(baseDir, 0755); err != nil {
		log.Fatalf("Error creating dashboard datastore directory: %s\n", err)
	}

	for _, board := range boards {
		dash, err := cli.GetDashboard(board.GetId())
		if err != nil {
			log.Fatalf("Error getting single dashboard: %s", err)
		}

		jsc, err := json.MarshalIndent(dash, "", "  ")
		if err != nil {
			log.Fatalf("Error unmarshaling responded JSON object: %s\n", err)
		}

		file, err := os.Create(filepath.Join(baseDir, toValidFileName(board.GetTitle())+".json"))
		if err != nil {
			log.Fatalf("Error creating json file for dashboard: %s\n", err)
		}
		file.Write(jsc)
		file.Close()
	}
}

// toValidFileName converts forbidden characters in UNIX/Linux file name to valid one.
// Strictly speaking, whitespace is allowed to be used for file name on UNIX/Linux machine but it makes hard to see.
// The original dashboard title would be remained in the exported data.
func toValidFileName(s string) string {
	repl := strings.NewReplacer(" ", "_", "/", "")
	return repl.Replace(s)
}

func init() {
	dashboardCmd.AddCommand(dashboardExportCmd)

	dashboardExportCmd.Flags().StringVarP(&targetDir, "--target-dir", "d", "",
		"already existing destination directory (default is current directory)")
}
