package main

import "fmt"

func main() {
	// Create anonymous struct
	point := struct {
		X, Y int
	}{
		X: 1, Y: 2,
	}

	// Access and print the struct fields
	fmt.Printf("Point coordinates: (%d, %d)\n", point.X, point.Y)

	// You can also modify the fields
	point.X = 10
	point.Y = 5
	fmt.Printf("Modified point: %+v\n", point)

	// Create another instance with shorthand syntax
	anotherPoint := struct {
		X, Y int
	}{1, 2}
	fmt.Printf("Another point: %+v\n", anotherPoint)
}
