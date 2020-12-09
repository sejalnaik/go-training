package main

import (
	"channels-task/customer"
	"fmt"
)

func main() {
	ch := make(chan int)
	customer1 := customer.NewCustomer(1, "sejal", "naik")
	go customer1.PlaceOrder(ch, 2, 10.0)
	go customer1.PlaceOrder(ch, 3, 10.0)
	go customer1.PlaceOrder(ch, 4, 10.0)
	go customer1.PlaceOrder(ch, 5, 10.0)
	fmt.Println("Order placed:", <-ch)
}
