package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func startServer() {

	initRoutes()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(listener, nil)
}

func initRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("index.html")
		if err != nil {
			log.Fatal(err)
		}
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(content)
	})
}
