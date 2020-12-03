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
	var p *int = &a
	fmt.Println("Value of p: ", p)
	//this gives us the value held by the address inside p
	fmt.Println("Deref of p: ", *p)
	fmt.Println("Type of p: ", reflect.TypeOf(p))
	fmt.Println("Memory address of p: ", &p)

	/*pointerToPointer is a variable of type **int since it is holding address of
	variable already holding address of another variable of type *int*/
	var pointerToPointer **int = &p
	fmt.Println("Value of pointerToPointer: ", pointerToPointer)
	//this gives us the value held by the address inside p
	fmt.Println("Deref of pointerToPointer: ", *pointerToPointer)
	//this gives us the value held by the address twice up
	fmt.Println("Double Deref of pointerToPointer: ", **pointerToPointer)
	fmt.Println("Type of pointerToPointer: ", reflect.TypeOf(pointerToPointer))
	fmt.Println("Memory address of pointerToPointer: ", &pointerToPointer)

}
