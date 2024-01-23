package main
import (
    "os"
    "log"
    "fmt"
    "net/http" 
)
var path = ""
func main() {
    //get current path
    path, err := os.Getwd()

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(path)

    fs := http.FileServer(http.Dir(path))
    log.Fatal(http.ListenAndServe(":8000",fs))
}