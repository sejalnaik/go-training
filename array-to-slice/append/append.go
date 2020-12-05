package main

import "fmt"

func main() {

	myArray := [5]int{10, 20, 30, 40, 50}

	//to remove 30 from the array
	sliceOne := myArray[0:2]
	fmt.Println(sliceOne)

	sliceTwo := myArray[3:]
	fmt.Println(sliceTwo)

	newSlice := append(sliceOne, sliceTwo...)
	fmt.Println(newSlice)

}
