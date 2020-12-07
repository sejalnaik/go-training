package main

import "fmt"

func main() {
	var myInt int = 42
	myIntPointer := &myInt
	fmt.Println(*myIntPointer)
}
