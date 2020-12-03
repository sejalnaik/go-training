package main

import (
	"fmt"
)

//package level varibales, not necessary to be used
var name string

//can only be assigned value if within the statement
var myName string = "sejal naik"

func main() {
	//declaring one variable
	var rollNo int
	rollNo = 20
	fmt.Println(rollNo)

	//declaring multiple variables at the same time
	var (
		firstName, lastName, flag = "sejal", "naik", true
	)
	fmt.Println(firstName, lastName, flag)

	//assigning value to package variable
	name = "sejal naik"
	fmt.Println(name)

	//printing value of already assigned package variable
	fmt.Println(myName)

	//shorthand type spcify and value assignment
	test := 10
	fmt.Printf("%T", test)
}
