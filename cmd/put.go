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
		fmt.Printf("Going to put [%s]: [%s] in KV-store\n", args[0], args[1])
		key := args[0]
		value := args[1]

		// Double-checking to not add empty value as it could break delete functionality
		if value == "" {
			log.Fatal("Value of key cannot be empty")
			os.Exit(1)
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

		if err != nil {
			panic(err)
		}

		fmt.Printf("Successfully put [%s]: [%s] pair", args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
