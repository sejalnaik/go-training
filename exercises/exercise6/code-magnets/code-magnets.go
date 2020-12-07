package main

import "fmt"

func main() {
	fmt.Println(sum(4, 2, 1))
	fmt.Println(sum(9, 7))
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
