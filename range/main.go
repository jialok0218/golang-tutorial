package main

import (
	"fmt"
	"time"
)

func main() {
	// Array to demonstrate range
	numbers := []int{1, 2, 3, 4, 5}

	// 1. Range with both index and value
	fmt.Println("Index and value:")
	for i, e := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, e)
	}

	// 2. Range with only value
	fmt.Println("\nOnly values:")
	for _, e := range numbers {
		fmt.Printf("Value: %d\n", e)
	}

	// 3. Range with only index
	fmt.Println("\nOnly indices:")
	for i := range numbers {
		fmt.Printf("Index: %d\n", i)
	}

	// 4. Range with timer (will run for 3 seconds)
	fmt.Println("\nTimer example:")
	count := 0
	for range time.Tick(time.Second) {
		fmt.Println("Tick")
		count++
		if count >= 3 {
			break
		}
	}
}
