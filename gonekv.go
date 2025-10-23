package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func contains(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func parseCommandLine() (string, []string, error) {

	if len(os.Args) <= 1 {
		return "", nil, errors.New("gonekv: Missing required arguments: gonekv [action] [operand(s)]\nExample Usage: gonekv set '1' 'Lorem ipsum'")
	}

	listOfActions := []string{"get", "set", "put", "delete"}
	if !contains(listOfActions, os.Args[1]) {
		return "", nil, fmt.Errorf("gonekv: Invalid action provided\nList of actions supported: %v", listOfActions)
	}

	action := os.Args[1]
	var errMsg string = ""
	var args []string
	switch action {
	case "get", "delete":
		if len(os.Args) <= 2 {
			errMsg = fmt.Sprintf("gonekv: Missing operand: gonekv [action] [key]\nExample Usage: gonekv %s '1'", action)
			break
		}

		args = append(args, os.Args[2]) // Get key text

	case "set", "put":
		if len(os.Args) <= 3 {
			errMsg = fmt.Sprintf("gonekv: Missing operand: gonekv [action] [key] [value]\nExample Usage: gonekv %s '1' 'Lorem ipsum'", action)
			break
		}

		args = append(args, os.Args[2]) // Get key text
		args = append(args, os.Args[3]) // Get value text

	default:
		log.Fatal("Invalid action found: " + action)
		os.Exit(1)
	}
	if errMsg != "" {
		return action, nil, errors.New(errMsg)
	}

	return action, args, nil
}

func handleGet(args []string) (string, error) {
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
		var currentLineValue = strings.Split(currentLine, ":")[1]

		if currentLineKey == key && currentLineValue != "" { // Ignore key if value is empty string i.e. deleted record
			value = strings.Split(currentLine, ":")[1]
		} // Find the last value for matched key
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return value, nil
}

func handleSet(args []string) error {
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

	return nil
}

func handlePut(args []string) error {

	// TODO

	return nil
}

func handleDelete(args []string) error {
	key := args[0]
	var pair = fmt.Sprintf("%s:\n", key)

	f, err := os.OpenFile("store.dat", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		panic("Failed to create/append to file: " + err.Error())
	}
	defer f.Close()

	if _, err = f.WriteString(pair); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	log.SetFlags(0) // Disable date time logging

	action, args, err := parseCommandLine()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("%s - %v\n", action, args)
	switch action {
	case "get":
		value, err := handleGet(args)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if value != "" {
			fmt.Printf("value: %s", value)
		} else {
			fmt.Printf("key [%s] does not exist", args[0])
		}

	case "set":
		err = handleSet(args)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Printf("Successfully set key-value pair => %s: %s", args[0], args[1])
	case "put":
		err = handlePut(args)
	case "delete":
		err = handleDelete(args)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
