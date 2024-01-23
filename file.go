package main

import (
	"fmt"
	"log"
	"os"
)
func saveFile(filename string,file_contents []byte){
	file,err := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil{
		log.Fatal(err)

	}
	
	defer file.Close()

	_, err = file.Write(file_contents)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success")

}

func rmFile(fileName string){
	err := os.Remove(fileName)

	// Check for errors
	if err != nil {
		// Handle the error
		panic(err)
	}
}
