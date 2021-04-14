package main

import (
	"fmt"
	"io/ioutil"
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
}

func main() {
	if fileExists(systemFilePath) {
		content, err := ioutil.ReadFile(systemFilePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", content)
	} else {
		log.Fatal("Error: Not a valid file")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
