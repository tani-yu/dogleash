package main

import "github.com/tani-yu/dogleash/cmd"
import _ "github.com/tani-yu/dogleash/cmd/dashboard"
import _ "github.com/tani-yu/dogleash/cmd/monitor"
import _ "github.com/tani-yu/dogleash/cmd/config"

func main() {
	cmd.Execute()
}
