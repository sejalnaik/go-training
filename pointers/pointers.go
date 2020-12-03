package main

import (
	"fmt"
	"reflect"
)

func main() {
	// a is a variable of type int
	var a int = 10
	fmt.Println("Value of a: ")
	fmt.Println("Memory address of a: ", &a)

	// p is a vraibale of type *int since it is holding address of varibale whose type is int
	p := &a
	fmt.Println("Value of p: ", p)
	//this gives us the value held by the address inside p
	fmt.Println("Deref of p: ", *p)
	fmt.Println("Type of p: ", reflect.TypeOf(p))
	fmt.Println("Memory address of p: ", &p)

	var hello string = "hello"

	ps := &hello
	fmt.Println("Value of p: ", ps)
	fmt.Println("Deref of p: ", *ps)
	fmt.Println("Type of p: ", reflect.TypeOf(ps))
	fmt.Println("Memory address of p: ", &ps)

}
