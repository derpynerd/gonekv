package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get [key] from KV-store",
	Long:  "Get [key] from KV-store if it has been set previously",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Going to get [%s] from KV-store\n", args[0])
		key := args[0]

		f, err := os.OpenFile("store.dat", os.O_RDONLY, 0664)
		if err != nil {
			panic("Failed to open/read file: " + err.Error())
		}
		defer f.Close()

		var value = ""
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var currentLine = scanner.Text()
			var currentLineKey = strings.Split(currentLine, ":")[0]

			if currentLineKey == key { // Ignore key if value is empty string i.e. deleted record
				value = strings.Split(currentLine, ":")[1]
			} // Find the last value for matched key
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if value != "" {
			fmt.Printf("Value for [%s]: %s", key, value)
		} else {
			fmt.Printf("Key [%s] does not exist", args[0]) // If value is empty string then it is a deleted record
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
