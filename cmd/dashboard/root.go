//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package dashboard

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tani-yu/dogleash/cmd"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

// dashboardCmd represents the dashboardCmd command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Perform operations related to dashboard of Datadog",
}

// getJSONDataFromAPI **WORKAROUND** get JSON from dashboard api
func getJSONDataFromAPI(path string) []byte {

	u, err := url.Parse("https://api.datadoghq.com/api/v1/" + path)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("api_key", viper.GetString("api_key"))
	q.Set("application_key", viper.GetString("app_key"))
	u.RawQuery = q.Encode()

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	jd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return jd
}

// GetAllDashboards **WORKAROUND** get all dashboard with widget form
func GetAllDashboards() []datadog.Board {
	jd := getJSONDataFromAPI("dashboard")

	var boards map[string][]datadog.Board
	err := json.Unmarshal([]byte(jd), &boards)
	if err != nil {
		log.Fatal(err)
	}

	return boards["dashboards"]
}

func init() {
	cmd.RootCmd.AddCommand(dashboardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// computeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// computeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
