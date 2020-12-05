package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3}
	for i, number := range mySlice {
		fmt.Println(i, number)
	}
}
