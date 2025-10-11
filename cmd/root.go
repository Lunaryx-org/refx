package cmd

import (
	"fmt"
	"os"

	"github.com/Lunaryx-org/refx/shared"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "refx [old-path] [new-path]",
	Short: "Replace Go import paths across your project",
	Long:  `refx replaces import paths in all .go files in the current directory`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		oldPath := args[0]
		newPath := args[1]

		if err := shared.Fileio(oldPath, newPath, verbose); err != nil {
			fmt.Fprintf(os.Stderr, "Error %s\n", err)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
