package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	p := Vertex{
		X: 1,
		Y: 2,
	}
	q := &p
	fmt.Printf("p: %+v\n", p)
	fmt.Printf("q points to: %+v\n", *q)
}
