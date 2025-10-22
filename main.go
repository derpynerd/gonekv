package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func contains(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func getActionAndArgs() (string, []string, error) {

	if len(os.Args) <= 1 {
		return "", nil, errors.New("gonekv: Missing required arguments: gonekv [action] [operand(s)]\nExample Usage: gonekv set '1' 'Lorem ipsum'")
	}

	listOfActions := []string{"get", "set", "put", "delete"}
	if !contains(listOfActions, os.Args[1]) {
		return "", nil, fmt.Errorf("gonekv: Invalid action provided\nList of actions supported: %v", listOfActions)
	}

	action := os.Args[1]
	var errMsg string = ""
	switch action {
	case "get", "delete":
		if len(os.Args) <= 2 {
			errMsg = fmt.Sprintf("gonekv: Missing operand: gonekv [action] [key]\nExample Usage: gonekv %s '1'", action)
		}

	case "set", "put":
		if len(os.Args) <= 3 {
			errMsg = fmt.Sprintf("gonekv: Missing operand: gonekv [action] [key] [value]\nExample Usage: gonekv %s '1' 'Lorem ipsum'", action)
		}

	default:
		panic("Invalid action found: " + action)
	}
	if errMsg != "" {
		return action, nil, errors.New(errMsg)
	}

	return action, nil, nil
}

func handleGet(args []string) error {

	return nil
}

func handleSet(args []string) error {

	// fd, err = os.WriteFile("store.dat")

	return nil
}

func handlePut(args []string) error {

	return nil
}

func handleDelete(args []string) error {

	return nil
}

func main() {
	log.SetFlags(0) // Disable date time logging

	action, args, err := getActionAndArgs()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("%s", action)
	switch action {
	case "get":
		err = handleGet(args)
	case "set":
		err = handleSet(args)
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
