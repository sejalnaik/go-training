package main

import "fmt"

func main() {
	myArray := [3]int{1, 2, 3}
	fmt.Println("myArray:", myArray)
	fmt.Println("Address of myArray:", &myArray[0])

	secondArray := myArray
	fmt.Println("secondArray:", secondArray)
	fmt.Println("Address of secondArray:", &secondArray[0])
}
