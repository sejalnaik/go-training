package main

import "fmt"

func main() {
	var person1 = person{
		firstName: "sejal",
		lastName:  "naik",
		age:       26,
	}

	fmt.Println("Before modifying person1:", person1)
	modifyPerson(person1)
	fmt.Println("After modifying person1:", person1)
}

func modifyPerson(p person) {
	p.firstName = "Sej"
	p.age = 50
	fmt.Println("In modifyPerson, person1:", p)
}

type person struct {
	firstName string
	lastName  string
	age       int
}
