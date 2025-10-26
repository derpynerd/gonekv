package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gonekv",
	Short: "gonekv is a simple key-value store",
	Long:  "gonekv is a simple key-value store which can be used as a mini-database for storage",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred while running gonekv")
		os.Exit(1)
	}
}
