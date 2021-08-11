package main

import (
	"fmt"
	"log"
	"os"
)

var systemFilePath string

func init() {
	if len(os.Args) > 1 {
		systemFilePath = os.Args[1]
	} else {
		log.Fatal("Error: No argument passed.")
	}
	if !fileExists(systemFilePath) {
		log.Fatal("Error: Not a valid file")
	}
}

func main() {
	content, err := os.ReadFile(systemFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
