package main

import "github.com/tani-yu/dogleash/cmd"
import _ "github.com/tani-yu/dogleash/cmd/dashboard"
import _ "github.com/tani-yu/dogleash/cmd/monitor"

func main() {
	cmd.Execute()
}
