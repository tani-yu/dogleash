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
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tani-yu/dogleash/cmd"
	datadog "gopkg.in/zorkian/go-datadog-api.v2"
)

// MyDashboards widget form dashboard struct
type MyDashboards struct {
	MyDashboards []datadog.Board `json:"dashboards,omitempty"`
}

// dashboardCmd represents the dashboardCmd command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Perform operations related to dashboard of Datadog",
}

// getJSONDataFromAPI **WORKAROUND** get JSON from dashboard api
func getJSONDataFromAPI(path string) []byte {
	baseurl := "https://api.datadoghq.com/api/v1/"
	param := "?api_key=" + viper.GetString("api_key") + "&application_key=" + viper.GetString("app_key")

	url := baseurl + path + param
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(url)
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

	var boards MyDashboards
	err := json.Unmarshal([]byte(jd), &boards)
	if err != nil {
		log.Fatal(err)
	}

	return boards.MyDashboards

	// err := cli.doJsonRequest("GET", "/v1/dashboard/", nil, &boards)
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
