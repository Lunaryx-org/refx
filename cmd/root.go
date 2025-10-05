package cmd

import (
	"fmt"
	"os"

	"github.com/Lunaryx-org/refx/shared"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "img [old-path] [new-path]",
	Short: "Replace Go import paths across your project",
	Long:  `img replaces import paths in all .go files in the current directory`,
	Args:  cobra.ExactArgs(2), // Requires exactly 2 arguments
	Run: func(cmd *cobra.Command, args []string) {
		oldPath := args[0]
		newPath := args[1]

		fmt.Println("old_path:	", oldPath)
		fmt.Println("new_path:	", newPath)
		// Call your actual logic here
		shared.Fileio(oldPath, newPath)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
