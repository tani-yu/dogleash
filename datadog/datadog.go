// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datadog

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

// NewDDClient returns the DataDog client
func NewDDClient() (*datadog.Client, error) {
	if viper.GetString("api_key") == "" {
		return nil, fmt.Errorf("Error: API key was not set. Please check your .dogrc file or DATADOG_API_KEY environment variable.")
	}
	if viper.GetString("app_key") == "" {
		return nil, fmt.Errorf("Error: Application key was not set. Please check your .dogrc file or DATADOG_APP_KEY environment variable.")
	}

	return datadog.NewClient(viper.GetString("api_key"), viper.GetString("app_key")), nil
}
