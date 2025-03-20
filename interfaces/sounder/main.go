package main

import "fmt"

// Base interfaces
type Sounder interface {
	MakeSound()
}

type Mover interface {
	Move()
}

// Combined interface using embedding
type Animal interface {
	Sounder // embeds Sounder interface
	Mover   // embeds Mover interface
}

// Dog struct
type Dog struct {
	Name string
}

// Dog implements both required methods
func (d Dog) MakeSound() {
	fmt.Printf("%s says: Woof!\n", d.Name)
}

func (d Dog) Move() {
	fmt.Printf("%s runs on four legs\n", d.Name)
}

func main() {
	// Dog must implement both MakeSound() and Move()
	dog := Dog{Name: "Buddy"}

	// Can use as Animal (both methods)
	var animal Animal = dog
	animal.MakeSound()
	animal.Move()

	// Can also use as just a Sounder
	var sounder Sounder = dog
	sounder.MakeSound()

	// Can also use as just a Mover
	var mover Mover = dog
	mover.Move()
}
