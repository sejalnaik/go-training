package main

import "fmt"

func main() {
	fmt.Println("start of main")

	//no return function
	displayName()

	//returnunt function
	userNo := addOne(10)
	fmt.Println(userNo)

	fmt.Println("end of main")
}

func displayName() {
	fmt.Println("Sejal")
}

func addOne(number int) int {
	number = number + 1
	return number
}
