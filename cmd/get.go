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

		value := Get(args[0])

		fmt.Printf("%s", value)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
