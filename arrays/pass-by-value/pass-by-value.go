package main

import "fmt"

func main() {
	myArray := [3]int{}
	fmt.Println("before change :", myArray)
	changeArray(myArray)
	fmt.Println("after change :", myArray)
	fmt.Println("address of array in main:", &myArray[0])
}

func changeArray(tempArray [3]int) {
	tempArray[0] = 1
	fmt.Println("address of array in changeArray:", &tempArray[0])
}
