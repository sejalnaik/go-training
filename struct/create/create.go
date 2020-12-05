package main

import "fmt"

func main() {
	var person1 = person{
		firstName: "sejal",
		age:       26,
	}

	fmt.Println("person1:", person1)
	fmt.Println("adress of person1:", &person1)
	fmt.Println("firstName:", person1.firstName)
	fmt.Println("lastName:", person1.lastName)
	fmt.Println("age:", person1.age)

	personAdress := &person1

	fmt.Println("firstName through address of person1:", personAdress.firstName)
}

type person struct {
	firstName string
	lastName  string
	age       int
}
