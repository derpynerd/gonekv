package main

import (
	"gonekv/cmd"
	"log"
)

func main() {
	log.SetFlags(0) // Disable date time logging

	cmd.Execute()
}
