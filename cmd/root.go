package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "sshm",
	Short: "A ssh program manager",
	Long:  `A ssh program manager,build with love by jiuzi in Go.`,
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
