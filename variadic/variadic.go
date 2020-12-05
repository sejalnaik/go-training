package main

import "fmt"

func main() {
	myArray := [5]int{10, 20, 30, 40, 50}
	displayArray(myArray[:]...)
}

func displayArray(numbers ...int) {
	fmt.Println(numbers)
}
