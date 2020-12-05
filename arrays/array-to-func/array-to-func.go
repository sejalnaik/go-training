package main

import "fmt"

func main() {
	myArray := [3]int{1, 2, 3}
	displayArrayOne(myArray)
	displayArrayTwo(1, 2, 3)
}

func displayArrayOne(tempArray [3]int) {
	fmt.Println(tempArray)
}

func displayArrayTwo(tempArray ...int) {
	fmt.Println(tempArray)
}
