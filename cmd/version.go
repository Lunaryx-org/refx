package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCMD = &cobra.Command{
	Use:   "version",
	Short: "print the version of the app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.3.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCMD)
}
