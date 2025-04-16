package main

import "fmt"

type Pizza struct {
	Name        string
	Size        string
	Toppings    []string
	IsDelicious bool
}

func pizzaStyle(p Pizza) string {
	return p.Name + " pizza is a " + p.Size + " pizza with toppings of: " + fmt.Sprint(p.Toppings)
}

type Restaurant struct {
	Name      string
	Rating    int
	PizzaMenu []Pizza
}

func restaurantInfo(r Restaurant) string {
	return r.Name + " has a rating of " + fmt.Sprint(r.Rating) + " and serves the following pizzas: " + fmt.Sprint(r.PizzaMenu)
}

func main() {
	myPizza := Pizza{
		Name:        "Margherita",
		Size:        "medium",
		Toppings:    []string{"tomatoes", "mozzarella", "basil"},
		IsDelicious: true,
	}

	myRestaurant := Restaurant{
		Name:      "Pizzeria del Corso",
		Rating:    4,
		PizzaMenu: []Pizza{myPizza},
	}

	fmt.Println(pizzaStyle(myPizza))
	fmt.Println(restaurantInfo(myRestaurant))
}
