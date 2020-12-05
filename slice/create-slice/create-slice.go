package main

import "fmt"

func main() {
	//methods of creating slice
	//1)
	var firstSlice []int
	firstSlice = make([]int, 3)
	firstSlice[0] = 1
	firstSlice[1] = 2
	firstSlice[2] = 3
	fmt.Println(firstSlice)

	//2)
	secondSlice := make([]int, 3)
	secondSlice[0] = 1
	secondSlice[1] = 2
	secondSlice[2] = 3
	fmt.Println(secondSlice)

	//3)
	thirdSlice := []int{1, 2, 3}
	fmt.Println(thirdSlice)
}
