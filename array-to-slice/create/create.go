package main

import "fmt"

func main() {
	myArray := [5]int{10, 20, 30, 40, 50}
	sliceOne := myArray[0:2]
	fmt.Println("myArray[0:2]:", sliceOne)

	sliceOne = append(sliceOne, 10, 20, 30)
	fmt.Println(lr)

	sliceTwo := myArray[2:2]
	fmt.Println("myArray[2:2]:", sliceTwo)

	//sliceThree := myArray[2:0]
	//fmt.Println("myArray[2:0]:", sliceThree)

	sliceFour := myArray[2:]
	fmt.Println("myArray[2:]:", sliceFour)

	sliceFive := myArray[:2]
	fmt.Println("myArray[:2]", sliceFive)

	sliceSix := myArray[:]
	fmt.Println("myArray[:]", sliceSix)
}
