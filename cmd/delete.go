package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete [key] pair from KV-store",
	Long:  "Delete [key] pair from KV-store permanently using Append-Only logic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Going to delete [%s] from KV-store\n", args[0])
		key := args[0]
		var pair = fmt.Sprintf("%s:\n", key) // Add new record for key with empty value - signifying that this record is deleted

		f, err := os.OpenFile("store.dat", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			panic("Failed to create/append to file: " + err.Error())
		}
		defer f.Close()

		if _, err = f.WriteString(pair); err != nil {
			panic(err)
		}

		fmt.Printf("Successfully deleted [%s] key", args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
