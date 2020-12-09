package customer

import (
	"channels-task/order"
	"math/rand"
)

// Customer to create a type for customer
type Customer struct {
	id        int
	firstName string
	lastName  string
	orders    []*order.Order
}

// NewCustomer to create struct of type customer
func NewCustomer(id int, firstName string, lastName string) *Customer {
	tempCustomer := &Customer{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		orders:    make([]*order.Order, 0),
	}
	return tempCustomer
}

// PlaceOrder to place order by the customer
func (c *Customer) PlaceOrder(ch chan int, quantity int, cost float64) {
	randomID := rand.Intn(100)
	tempOrder := order.NewOrder(randomID, quantity, cost)
	c.orders = append(c.orders, tempOrder)
	ch <- randomID
}
