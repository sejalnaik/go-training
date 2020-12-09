package order

// Order to create a type for order
type Order struct {
	id       uint
	quantity int
	cost     float64
}

// NewOrder to create struct of type order
func NewOrder(id uint, quantity int, cost float64) *Order {
	tempOrder := &Order{
		id:       id,
		quantity: quantity,
		cost:     cost,
	}
	return tempOrder
}
