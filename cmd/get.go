package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get [key] from KV-store",
	Long:  "Get [key] from KV-store if it has been set previously",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Going to get [%s] from KV-store\n", args[0])

		if !Get(args[0]) {
			fmt.Printf("Failed to get key [%s] from KV-store", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
