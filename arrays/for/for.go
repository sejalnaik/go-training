package main

import "fmt"

func main() {
	myArray := [3]int{1, 2, 3}
	for i, number := range myArray {
		fmt.Println(i, number)
	}
}
