package main

import "fmt"

func main() {
	race(&car{name: "nano"})
	race(&bus{name: "best"})
	race(&bike{name: "bullet"})
}

type car struct {
	name string
}
type bus struct {
	name string
}
type bike struct {
	name string
}

type movable interface {
	move() string
}

func (c *car) move() string {
	return c.name + " is moving"
}

func (b *bus) move() string {
	return b.name + " is moving"
}

func (b *bike) move() string {
	return b.name + " is moving"
}

func race(mov movable) {
	fmt.Println(mov.move())
}
