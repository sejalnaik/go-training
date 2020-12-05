package main

import "fmt"

func main() {
	mySlice := make([]int, 3, 4)
	fmt.Println("Before append")
	fmt.Println("Slice:", mySlice)
	fmt.Println("Adress of mySlice:", &mySlice[0])
	fmt.Println("Length of slice:", len(mySlice))
	fmt.Println("Capacity of slice:", cap(mySlice))

	mySlice = append(mySlice, 4)

	fmt.Println("After append")
	fmt.Println("Slice:", mySlice)
	fmt.Println("Adress of mySlice:", &mySlice[0])
	fmt.Println("Length of slice:", len(mySlice))
	fmt.Println("Capacity of slice:", cap(mySlice))

	mySlice = append(mySlice, 5, 6)

	fmt.Println("After append")
	fmt.Println("Slice:", mySlice)
	fmt.Println("Adress of mySlice:", &mySlice[0])
	fmt.Println("Length of slice:", len(mySlice))
	fmt.Println("Capacity of slice:", cap(mySlice))

}
