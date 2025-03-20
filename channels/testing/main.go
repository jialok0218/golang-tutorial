package main

import "fmt"

func main() {

	// create channel
	ch := make(chan string)

	// function call with goroutine
	go receiveData(ch)

	//receive channel data
	fmt.Println(<-ch)

}

func receiveData(ch chan string) {

	// receive data from the channel
	ch <- "Cool do you see that"

}
