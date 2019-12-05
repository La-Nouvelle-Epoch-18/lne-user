package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/La-Nouvelle-Epoch-18/lne-user/cmd/start"
	"github.com/La-Nouvelle-Epoch-18/lne-user/cmd/version"
)

func init() {
	root.AddCommand(version.Cmd)
	root.AddCommand(start.Cmd)
}

var root = &cobra.Command{
	Use:   "lne-user",
	Short: "lne-user",
}

// main command line
func main() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
