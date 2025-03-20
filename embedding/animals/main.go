package main

import "fmt"

// Define interface for animals
type Speaker interface {
	Speak() string
}

type Animal struct {
	Name string
}

type Dog struct {
	Animal
	Breed string
}

// Method for Animal - now returns string to implement Speaker interface
func (a Animal) Speak() string {
	return fmt.Sprintf("I am %s", a.Name)
}

// Dog-specific Speak method
func (d Dog) Speak() string {
	return fmt.Sprintf("Woof! %s and I'm a %s", d.Animal.Speak(), d.Breed)
}

// Function that works with any Speaker
func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	// Create a dog
	dog := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "German Shepherd",
	}

	// Create regular animal
	animal := Animal{Name: "Generic Animal"}

	// Use interface
	MakeSpeak(dog)    // Will use Dog's Speak method
	MakeSpeak(animal) // Will use Animal's Speak method
}
