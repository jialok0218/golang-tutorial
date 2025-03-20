package main

import "fmt"

func main() {
	// Basic slice declaration and initialization
	var a []int
	fmt.Printf("Empty slice a: %v\n", a)

	// Initialize with values
	a = []int{1, 2, 3, 4}
	fmt.Printf("Initialized slice a: %v\n", a)

	// Shorthand declaration
	b := []int{1, 2, 3, 4}
	fmt.Printf("Slice b: %v\n", b)

	// Slice with specific index initialization
	chars := []string{0: "a", 2: "c", 1: "b"}
	fmt.Printf("Slice chars: %v\n", chars)

	// Slicing operations
	c := b[1:4]
	fmt.Printf("Slice b[1:4]: %v\n", c)

	d := b[:3]
	fmt.Printf("Slice b[:3]: %v\n", d)

	e := b[3:]
	fmt.Printf("Slice b[3:]: %v\n", e)

	// Append operations
	a = append(a, 17, 3)
	fmt.Printf("After append to a: %v\n", a)

	// Concatenate slices
	combined := append(a, b...)
	fmt.Printf("Combined slices: %v\n", combined)

	// Create slice with make
	f := make([]byte, 5, 5)
	fmt.Printf("Slice created with make: %v\n", f)

	// Create slice from array
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:]
	fmt.Printf("Slices from array: %v\n", s)
}
