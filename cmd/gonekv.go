package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
