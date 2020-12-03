package main

import "fmt"

func main() {
	cgpa := 6
	switch cgpa {
	case 5, 6:
		fmt.Println("5")
	case 10:
		fmt.Println("10")
	default:
		fmt.Println("any other cgpa")
	}
}
