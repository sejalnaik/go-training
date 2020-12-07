package main

import "fmt"

func main() {
	var i interface{} = "hello"
	//any type of value can be stored in i
	i = 10
	fmt.Println(i)

	i = "hello"
	//check if type of i is string
	storeString, ok := i.(string)
	fmt.Println(storeString, ok)

	//check if type of i is int
	storeInt, ok := i.(int)
	fmt.Println(storeInt, ok)

	checkType("hello")
	checkType(10)
	checkType(10.0)

	var num float32 = 17
	checkType(num)

	checkType(true)
}

func checkType(anyType interface{}) {
	switch anyType.(type) {
	case string:
		fmt.Println("Argument is string")
	case float32:
		fmt.Println("Argument is float32")
	case float64:
		fmt.Println("Argument is float64")
	case int:
		fmt.Println("Argument is int")
	default:
		fmt.Println("other than my cases")
	}
}
