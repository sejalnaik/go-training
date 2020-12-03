package main

import "fmt"

func main() {
	fmt.Println(mathOperations(10, 20, add))
	fmt.Println(mathOperations(20, 10, subtract))
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func subtract(num1 int, num2 int) int {
	return num1 - num2
}

func mathOperations(number1 int, number2 int, func1 func(num1, num2 int) int) int {
	return func1(number1, number2)
}
