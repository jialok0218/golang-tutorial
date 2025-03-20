package main

import "fmt"

// Base struct (like a superclass)
type Animal struct {
	Name string
}

// Method for Animal
func (a Animal) MakeSound() {
	fmt.Println("Some generic animal sound")
}

// Dog struct embeds Animal (Composition)
type Dog struct {
	Animal // Embedding instead of inheritance
}

// Overriding behavior (method shadowing)
func (d Dog) MakeSound() {
	fmt.Printf("%s says: Woof woof!\n", d.Name)
}

func main() {
	dog := Dog{Animal{Name: "Buddy"}}
	dog.MakeSound() // Output: Buddy says: Woof woof!
}

/*
// Example Java Class with Inheritance
// Superclass
class Animal {
    String name;

    Animal(String name) {
        this.name = name;
    }

    void makeSound() {
        System.out.println("Some generic animal sound");
    }
}

// Subclass inheriting from Animal
class Dog extends Animal {
    Dog(String name) {
        super(name);
    }

    @Override
    void makeSound() {
        System.out.println(name + " says: Woof woof!");
    }
}

public class Main {
    public static void main(String[] args) {
        Dog dog = new Dog("Buddy");
        dog.makeSound(); // Output: Buddy says: Woof woof!
    }
}

*/
