package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Store [key]: [value] pair in KV-store",
	Long:  "Store [key]: [value] pair in KV-store which can be retrieved later",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		if Set(args[0], args[1]) {
			fmt.Printf("[%s]: [%s]", args[0], args[1])
		} else {
			fmt.Printf("Error: Failed to set [%s]: [%s] pair", args[0], args[1])
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
