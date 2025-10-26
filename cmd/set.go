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
		fmt.Printf("Going to set [%s]: [%s] in KV-store\n", args[0], args[1])

		if Set(args[0], args[1]) {
			fmt.Printf("Successfully set [%s]: [%s] pair", args[0], args[1])
		} else {
			fmt.Printf("Failed to set [%s]: [%s] pair", args[0], args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
