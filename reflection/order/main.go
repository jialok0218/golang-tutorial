package main

import (
	"fmt"
)

type order struct {
	orderId    int
	customerId int
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.orderId, o.customerId)
	return i
}

func main() {
	o := order{
		orderId:    1234,
		customerId: 567,
	}
	fmt.Println(createQuery(o))
}
