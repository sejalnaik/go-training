package main

import "fmt"

func main() {
	result := abc(func1)
	resultValue := result()
	fmt.Println(resultValue)
}

func result() string {
	resultValue := "This is result funtcion returned from abc"
	return resultValue
}

func func1(num1 int) int {
	return num1 + 1
}

func abc(func1 func(num1 int) int) func() string {
	number := func1(5)
	fmt.Println(number)
	return result
}
