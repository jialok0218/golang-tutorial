package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hi my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	person1 := Person{
		Name: "Alibaba",
		Age:  38,
	}

	person1.Introduce()
}
