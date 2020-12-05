package main

import "fmt"

func main() {
	mySlice := []int{0, 0, 0}
	fmt.Println("Before change :", mySlice)
	fmt.Println("Address of mySlice:", &mySlice[0])

	changeSlice(&mySlice)
	fmt.Println("After change :", mySlice)
	fmt.Println("Address of array in main:", &mySlice[0])
}

func changeSlice(tempSlice *[]int) {
	fmt.Println("In changeSlice, address of tempSlice", &(*tempSlice)[0])
	(*tempSlice)[0] = 1
}
