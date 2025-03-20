package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Basic switch with direct value

	operatingSystem := "darwin"
	switch operatingSystem {
	case "darwin":
		fmt.Println("Mac OS Hipster")
	case "linux":
		fmt.Println("Linux Geek")
	default:
		fmt.Println("Other")
	}

	// Switch with initialization statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	default:
		fmt.Printf("Running on %s\n", os)
	}

	// Switch with comparisons
	var number int = 10
	switch {
	case number < 10:
		fmt.Println("Smaller")
	case number == 10:
		fmt.Println("Equal")
	case number > 10:
		fmt.Println("Greater")
	default: // Optional
		fmt.Println("Default case")
	}

	// Switch with multiple cases
	var char byte = '?'
	switch char {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("Should escape")
	}
}
