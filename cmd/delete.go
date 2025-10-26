package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete [key] pair from KV-store",
	Long:    "Delete [key] pair from KV-store permanently using Append-Only logic",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if !Delete(args[0]) {
			fmt.Printf("Failed to delete key [%s]", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
