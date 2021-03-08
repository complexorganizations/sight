package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
