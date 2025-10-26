package cmd

import (
	"fmt"
	"log"
	"os"

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
		}

	},
}

func Put(key string, value string) (success bool) {

	existingValue := Get(key)

	// Checking if key exists, if not then exit
	if existingValue == "" {
		fmt.Printf("Key [%s] doesn't exist", key)
		return false
	}

	// Double-checking to not add empty value as it breaks delete functionality
	if value == "" {
		log.Fatal("Value of key cannot be empty")
		return false
	}
	var pair = fmt.Sprintf("%s:%s\n", key, value)

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
	rootCmd.AddCommand(putCmd)
}
