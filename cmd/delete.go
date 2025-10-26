package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete [key] pair from KV-store",
	Long:    "Delete [key] pair from KV-store permanently using Append-Only logic",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		Delete(args[0])

	},
}

func Delete(key string) (sucess bool) {

	existingValue := Get(key)

	// Checking if key exists, if not then exit
	if existingValue == "" {
		fmt.Printf("Key [%s] doesn't exist\n", key)
		return false
	}

	var pair = fmt.Sprintf("%s:\n", key) // Add new record for key with empty value - signifying that this record is deleted

	f, err := os.OpenFile("store.dat", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		panic("Failed to create/append to file: " + err.Error())
	}
	defer f.Close()

	if _, err = f.WriteString(pair); err != nil {
		panic(err)
	}

	return true
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
