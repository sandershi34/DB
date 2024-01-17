package main
import (
    "os"
    "log"
    "fmt"
)
var path = ""
func main() {
    //get current path
    path, err := os.Getwd()

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(path)
    listDirectories(path)
    path = changeDirectory(path,"newFolder")
    listDirectories(path)
    b := []byte("Hello world!")
    saveFile("helloworld.txt",b)
}