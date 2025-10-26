package cmd

import (
	"bufio"
	"fmt"
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

		value := Get(args[0])

		fmt.Printf("%s", value)

	},
}

func Get(key string) (success string) {

	f, err := os.OpenFile("store.dat", os.O_RDONLY, 0664)
	if err != nil {
		panic("Fatal: Failed to open/read file: " + err.Error())
	}
	defer f.Close()

	var value = "" // If value is empty string then it is a deleted record
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

	return value
}

func init() {
	rootCmd.AddCommand(getCmd)
}
