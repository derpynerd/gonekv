package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete [key] pair from KV-store",
	Long:  "Delete [key] pair from KV-store permanently using Append-Only logic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Going to delete [%s] from KV-store\n", args[0])

		if Delete(args[0]) {
			fmt.Printf("Successfully deleted key [%s]", args[0])
		} else {
			fmt.Printf("Failed to delete key [%s]", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
