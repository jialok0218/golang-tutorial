package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(message string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	for i := 1; i <= 5; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // Add one task to the WaitGroup

	go printMessage("Hello from Goroutine", &wg) // Starts a goroutine

	printMessage("Hello from Main", nil) // Runs in the main thread

	wg.Wait() // Wait for the goroutine to finish
}
