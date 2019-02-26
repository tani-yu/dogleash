//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// See Copyright Notice in LICENSE
//

package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	completionShells = map[string]func(w io.Writer) error{
		"bash": RootCmd.GenBashCompletion,
		"zsh":  RootCmd.GenZshCompletion,
	}
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: fmt.Sprintf("Output shell completion code for the specified shell (Supported: %s)", supportedShellNames()),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("Shell type not specified. (Supported: %s)", supportedShellNames()))
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Too many arguments. Expected only the shell type.")
			os.Exit(1)
		}

		run, found := completionShells[args[0]]
		if !found {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("Unsupported shell type %q. (Supported: %s)", args[0], supportedShellNames()))
			os.Exit(1)
		}

		if err := run(os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	},
}

// supportedShallNames returns string which contains names of the supported shell by DogLeash.
// The return value is generated from keys of completionShells map.
func supportedShellNames() string {
	shells := make([]string, 0)
	for s := range completionShells {
		shells = append(shells, s)
	}
	return strings.Join(shells, ", ")
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
