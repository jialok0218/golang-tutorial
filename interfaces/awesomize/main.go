package main

import "fmt"

// Interface declaration
type Awesomizer interface {
	Awesomize() string
}

// Struct type
type Foo struct{}

// Foo implements Awesomizer interface implicitly
func (foo Foo) Awesomize() string {
	return "Awesome!"
}

func main() {
	// Create an instance of Foo
	f := Foo{}

	// Use Foo directly
	fmt.Println("Direct call:", f.Awesomize())

	// Use Foo through interface
	var a Awesomizer = f
	fmt.Println("Through interface:", a.Awesomize())
}
