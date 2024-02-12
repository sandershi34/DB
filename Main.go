package main

import (
	"fmt"
	"log"
	"os"
)

var path = ""

func main() {
	//get current path
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
	startServer()
}
