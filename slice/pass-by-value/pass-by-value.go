package main

import "fmt"

func main() {
	mySlice := []int{0, 0, 0}
	fmt.Println("before change :", mySlice)
	changeSlice(mySlice)
	fmt.Println("after change :", mySlice)
	fmt.Println("address of array in main:", &mySlice[0])
}

func changeSlice(tempSlice []int) {
	tempSlice[0] = 1
	fmt.Println("address of array in changeSlice:", &tempSlice[0])
}
