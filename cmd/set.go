package cmd

import (
	"fmt"
	"log"
	"os"

	"gonekv/common"

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
		}

	},
}

func Set(key string, value string) (success bool) {

	existingValue := Get(key)

	// Checking if key exists, if not then exit
	if existingValue != "" {
		fmt.Printf("Key [%s] already exists", key)
		return false
	}

	// Double-checking to not add empty value as it breaks delete functionality
	if value == "" {
		log.Fatal("Value cannot be empty")
		return false
	}

	common.CompactionCheck()

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
	rootCmd.AddCommand(setCmd)
}
