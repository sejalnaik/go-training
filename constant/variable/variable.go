package main

import "fmt"

const myConstantString = "Hello"
const myConstantInt = 10

func main() {
	//private constant has camel casing

	fmt.Println("Value of my constant string is: ", myConstantString)

	fmt.Println("Value of my constant int is: ", myConstantInt)

	//shadowing a constant(new constant created)
	const myConstantString = "Hello world"
	fmt.Println("Value of my constant string is: ", myConstantString)

	//iota starts form 0 for every block
	const (
		_     = iota //first value of iota is 0
		cat   = iota
		dog   = iota
		camel = iota
		horse = iota
	)

	fmt.Println("Value of constants:", cat, dog, camel, horse)
}
