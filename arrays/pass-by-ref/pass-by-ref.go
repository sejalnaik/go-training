package main

import "fmt"

func main() {
	myArray := [3]int{}
	fmt.Println("Before change :", myArray)
	fmt.Println("Address of myArray:", &myArray[0])

	changeArray(&myArray)
	fmt.Println("After change :", myArray)
	fmt.Println("Address of array in main:", &myArray[0])
}

func changeArray(tempArray *[3]int) {
	fmt.Println("In changeArray, address of tempArray", &(*tempArray)[0])
	(*tempArray)[0] = 1
}
