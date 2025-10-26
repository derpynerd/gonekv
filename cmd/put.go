package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Update [key]: [value] pair in KV-store",
	Long:  "Update an existing [key]: [value] pair using Append-Only logic",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		if Put(args[0], args[1]) {
			fmt.Printf("[%s]: [%s]", args[0], args[1])
		} else {
			fmt.Printf("Error: Failed to set [%s]: [%s] pair", args[0], args[1])
		}

	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
