package main

import (
	"fmt"
	"time"
)

// Function to send data to a channel
func sendData(ch chan int) {
	ch <- 42 // Send a value to the channel
}

// Function to demonstrate select with channels
func doStuff(channelOut, channelIn chan int) {
	select {
	case channelOut <- 42:
		fmt.Println("We could write to channelOut!")
	case x := <-channelIn:
		fmt.Printf("We could read from channelIn: %d\n", x)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}

func main() {
	// 1. Unbuffered channel example
	fmt.Println("Unbuffered channel example:")
	ch := make(chan int) // Create an unbuffered channel for integers

	go sendData(ch) // Start a goroutine to send data
	v := <-ch       // Receive the value from the channel
	fmt.Println("Received:", v)

	// 2. Buffered channel example
	fmt.Println("\nBuffered channel example:")
	bufferedCh := make(chan int, 2) // Create a buffered channel with capacity 2
	bufferedCh <- 1                 // Send value 1
	bufferedCh <- 2                 // Send value 2
	close(bufferedCh)               // Close the channel

	// Read from the buffered channel until it is closed
	for i := range bufferedCh {
		fmt.Println("Read from buffered channel:", i)
	}

	// 3. Select statement example
	fmt.Println("\nSelect statement example:")
	channelOut := make(chan int)
	channelIn := make(chan int)

	go func() {
		time.Sleep(500 * time.Millisecond)
		channelIn <- 99 // Simulate sending a value to channelIn
	}()

	doStuff(channelOut, channelIn)

	// 4. Non-buffered channel with close example
	fmt.Println("\nNon-buffered channel with close example:")
	nonBufferedCh := make(chan int)
	go func() {
		nonBufferedCh <- 7 // Send a value
		close(nonBufferedCh)
	}()

	// Read from the channel and check if it is closed
	for v, ok := <-nonBufferedCh; ok; v, ok = <-nonBufferedCh {
		fmt.Println("Read from non-buffered channel:", v)
	}
}
