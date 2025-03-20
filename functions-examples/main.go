package main

import "fmt"

// 1. Function with no parameters and no return
func sayHello() {
	fmt.Println("Hello!")
}

// 2. Function with parameters but no return
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 3. Function with single return value
func add(a, b int) int {
	return a + b
}

// 4. Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 5. Function with named return values
func getCoordinates() (x, y int) {
	x = 10
	y = 20
	return // naked return
}

func main() {
	// Test the functions
	sayHello()

	greet("Alice")

	sum := add(5, 3)
	fmt.Printf("Sum: %d\n", sum)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Division result: %.2f\n", result)
	}

	x, y := getCoordinates()
	fmt.Printf("Coordinates: %d, %d\n", x, y)
}
