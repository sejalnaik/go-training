package main

import "fmt"

func main() {
	var userNo int

test:
	fmt.Println("Enter a number")
	fmt.Scanln(&userNo)
	fmt.Println("Your number is:", userNo)

	if userNo > 100 {
		fmt.Println("number too high")
		goto test
	}

	fmt.Println("Good number")
}
