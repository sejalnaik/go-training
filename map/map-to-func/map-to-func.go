package main

import "fmt"

func main() {
	myMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	newMap := changeMap(myMap)
	fmt.Println("newMap:", newMap)
	fmt.Println("myMap:", myMap)

}

func changeMap(newMap map[int]string) map[int]string {
	newMap[1] = "changed"
	return newMap
}
