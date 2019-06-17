package main

import "github.com/tani-yu/dogleash/cmd"
import _ "github.com/tani-yu/dogleash/cmd/dashboard"
import _ "github.com/tani-yu/dogleash/cmd/monitor"
import _ "github.com/tani-yu/dogleash/cmd/config"
import _ "github.com/tani-yu/dogleash/cmd/synthetics"
import _ "github.com/tani-yu/dogleash/cmd/synthetics/api"
import _ "github.com/tani-yu/dogleash/cmd/synthetics/browser"

func main() {
	cmd.Execute()
}
