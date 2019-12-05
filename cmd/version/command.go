package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
)

var Cmd = &cobra.Command{
	Use: "version",
	Run: printVersion,
}

func printVersion(*cobra.Command, []string) {
	fmt.Println(version)
}
