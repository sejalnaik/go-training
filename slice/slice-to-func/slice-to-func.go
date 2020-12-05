package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3}
	displaySliceOne(mySlice)
	displaySliceTwo(mySlice, mySlice, mySlice)
}

func displaySliceOne(tempSlice []int) {
	fmt.Println(tempSlice)
}

func displaySliceTwo(tempSlice ...[]int) {
	fmt.Println(tempSlice)
}
