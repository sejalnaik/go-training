package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return 22.0 / 7.0 * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * 22 / 7 * c.radius
}

func displayShape(sh shape) {
	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())
}

func main() {
	displayShape(&circle{radius: 10.0})
}
