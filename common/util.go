package common

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const compactionLimit = 5

func CompactionCheck() {

	file, err := os.OpenFile("store.dat", os.O_RDONLY, 0664)
	if err != nil {
		panic("Fatal: Failed to open/read file: " + err.Error())
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if lineCount >= compactionLimit {
		HandleCompaction(scanner, file)
	}

}

func HandleCompaction(scanner *bufio.Scanner, file *os.File) {

	_, err := file.Seek(0, io.SeekStart) // Seek to last line
	if err != nil {
		log.Fatal(err)
	}

	var lineCount = 0
	bucket := make(map[string][]int)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var currentLineKey = strings.Split(scanner.Text(), ":")[0]

		bucket[currentLineKey] = append(bucket[currentLineKey], lineCount)
		lineCount++

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Logging bucket: key -> value array
	for key, values := range bucket {
		strValues := strings.Trim(strings.Replace(fmt.Sprint(values), " ", ", ", -1), "[]")
		fmt.Printf("[%s] -> {%s}\n", key, strValues)
	}

	CleanupLines(bucket, file)

}

func CleanupLines(bucket map[string][]int, file *os.File) {

	// TODO: if key.value.list.count > 1 then delete all line indexes except last

}
