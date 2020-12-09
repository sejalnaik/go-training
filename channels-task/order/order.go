package order

// Order to create a type for order
type Order struct {
	id       int
	quantity int
	cost     float64
}

// NewOrder to create struct of type order
func NewOrder(id int, quantity int, cost float64) *Order {
	tempOrder := &Order{
		id:       id,
		quantity: quantity,
		cost:     cost,
	}
	return tempOrder
}
