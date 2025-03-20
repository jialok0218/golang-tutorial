package main

import "fmt"

// Vertex represents a coordinate point
type Vertex struct {
	X, Y float64
}

func main() {
	// Basic map operations
	m := make(map[string]int)
	m["key"] = 42
	fmt.Println("Value:", m["key"])

	// Delete and check existence
	delete(m, "key")
	value, ok := m["key"] // test if key "key" is present and retrieve it, if so
	fmt.Printf("Key exists: %v, Value: %d\n\n", ok, value)

	// Map literal with Vertex type
	locations := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	// Iterate over map content
	for key, value := range locations {
		fmt.Printf("Location: %s\n", key)
		fmt.Printf("Coordinates: %.5f, %.5f\n", value.X, value.Y)
		fmt.Println()
	}
}
