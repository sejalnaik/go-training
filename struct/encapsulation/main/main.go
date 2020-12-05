package main

import (
	"fmt"
	"struct/encapsulation/personstruct"
)

func main() {
	var person1 = personstruct.NewPerson("sejal", "naik", 26)

	fmt.Println("person1:", person1)
	fmt.Println("adress of person1:", &person1)
	fmt.Println("firstName:", person1.FirstName)
	fmt.Println("lastName:", person1.LastName)
	fmt.Println("age:", person1.Age)

}
