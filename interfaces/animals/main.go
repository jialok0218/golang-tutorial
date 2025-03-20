package main

import "fmt"

// Defining the Animal interface
type Animal interface {
	MakeSound() // Any type implementing this method is an Animal
}

// Dog struct
type Dog struct {
	Name string
}

// Dog implements the MakeSound method
func (d Dog) MakeSound() {
	fmt.Printf("%s says: Woof woof!\n", d.Name)
}

// Cat struct
type Cat struct {
	Name string
}

// Cat implements the MakeSound method
func (c Cat) MakeSound() {
	fmt.Printf("%s says: Meow!\n", c.Name)
}

func main() {
	// Polymorphism: An array of Animal interface types
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Kitty"},
	}

	// Loop through the animals and call MakeSound()
	for _, animal := range animals {
		animal.MakeSound() // Each type calls its own MakeSound method
	}
}
