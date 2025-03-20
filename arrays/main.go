package main

import (
	"fmt"
)

func main() {
	// Declare array of size 10
	var a [5]int
	a[2] = 42
	a[3] = 100
	i := a[3]
	fmt.Printf("Array a[3]: %d\n", i)

	// Different ways to declare and initialize arrays
	var b = [2]int{1, 2}
	fmt.Printf("Array b: %v\n", b)

	c := [2]int{1, 2}
	fmt.Printf("Array c: %v\n", c)

	d := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Array d: %v\n", d)
	fmt.Printf("Length of d: %d\n", len(d))
}
