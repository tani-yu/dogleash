// Copyright © 2018 The DogLeash Authors.
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

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version represents the current version of DogLeash CLI tool
const Version = "0.2.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  "Print the version of DogLeash CLI tool and its build information",
	Run:   printVersion,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func printVersion(cmd *cobra.Command, args []string) {
	fmt.Println("DogLeash Version:", Version)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Printf("Go OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
