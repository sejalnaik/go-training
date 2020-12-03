package main

import "fmt"

func main() {
	number := 10
	passByValue(number)
	fmt.Println("After calling passByValue, in main the value of number: ", number)
	fmt.Println("After calling passByValue, in main the address of number: ", &number)
}

func passByValue(number int) {
	number = number + 2
	fmt.Println("In passByValue func the value of number: ", number)
	fmt.Println("In passByValue func the address of number: ", &number)
}
