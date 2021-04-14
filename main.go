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
	content, err := ioutil.ReadFile(systemFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
