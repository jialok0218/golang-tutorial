package main

import (
	"log"
	"os"
)

// Define interfaces
type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

// Combined interface through embedding
type ReadWriter interface {
	Reader
	Writer
}

// Server struct with embedded Logger
type Server struct {
	Host string
	Port int
	*log.Logger
}

func main() {
	// Initialize server with embedded logger
	server := &Server{
		Host:   "localhost",
		Port:   80,
		Logger: log.New(os.Stdout, "SERVER: ", log.Ltime),
	}

	// Use embedded logger methods
	server.Println("Server starting...") // Calls server.Logger.Println

	// Access embedded logger directly
	var logger *log.Logger = server.Logger
	logger.Println("Direct logger access")

	// Access other fields
	server.Printf("Running on %s:%d\n", server.Host, server.Port)
}
