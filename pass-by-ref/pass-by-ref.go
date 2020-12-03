package main

import "fmt"

func main() {
	number := 10
	passByRef(&number)
	fmt.Println("After calling passByRef, in main the value of number: ", number)
	fmt.Println("After calling passByRef, in main the address of number: ", &number)
}

func passByRef(number *int) {
	*number = *number + 2
	fmt.Println("In passByRef func the value of number: ", number)
	fmt.Println("In passByRef func the address of number: ", &number)
}
