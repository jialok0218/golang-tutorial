package main

import "fmt"

func main() {
	// Creating vertices
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1, Y: 2}
	vertices := []Vertex{{1, 2}, {5, 2}, {5, 5}}

	// Accessing and modifying members
	v1.X = 4

	// Using methods
	fmt.Printf("v1 absolute distance: %.2f\n", v1.Abs())

	// Using pointer method
	fmt.Printf("v2 before add: %+v\n", v2)
	v2.add(3)
	fmt.Printf("v2 after add: %+v\n", v2)

	// Print slice of vertices
	fmt.Println("\nAll vertices:")
	for i, v := range vertices {
		fmt.Printf("vertex %d: %+v\n", i, v)
	}
}
