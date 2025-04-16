package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

// content holds the static content (2 files) for the web server.
//
//go:embed a.txt b.txt
var content embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(content)))
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
