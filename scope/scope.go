package main

import "fmt"

// package level variable
var num1 int = 10

func main() {
	fmt.Println("package variable num1:", num1)
	displayNum()
}

func displayNum() {
	//local variable wuth same name presides over package level variable
	var num1 int = 20
	fmt.Println("function variable num1:", num1)
}
