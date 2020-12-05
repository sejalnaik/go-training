package main

import "fmt"

func main() {
	var myMap = make(map[string]int)

	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 4
	myMap["three"] = 3
	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	delete(myMap, "three")
	fmt.Println(myMap)

	var num1 = myMap["one"]
	fmt.Println(num1)

	var num2, ok = myMap["twoo"]
	fmt.Println(num2, ok)

	var num3, ok2 = myMap["one"]
	fmt.Println(num3, ok2)
}
