package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Get(key string) (success bool) {

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

	if value != "" {
		fmt.Printf("Value for [%s]: %s", key, value)
	} else {
		fmt.Printf("Key [%s] does not exist", key) // If value is empty string then it is a deleted record
	}

	return true
}

func Set(key string, value string) (success bool) {

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

	return true
}

func Put(key string, value string) (success bool) {

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

	return true
}

func Delete(key string) (sucess bool) {

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
