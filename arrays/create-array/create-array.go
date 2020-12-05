package main

import (
	"fmt"
)

func main() {
	// methods of creating array
	//1)
	var firstArray [3]int
	fmt.Println(firstArray)

	//2)
	secondArray := [3]int{}
	fmt.Println(secondArray)

	//3)
	thirdArray := [3]int{1, 2, 3}
	fmt.Println(thirdArray)

}
