package main

import "fmt"

func main() {
	var person1 = newPerson("sejal", "naik", 26)

	fmt.Println("person1:", person1)
	fmt.Println("adress of person1:", &person1)
	fmt.Println("firstName:", person1.firstName)
	fmt.Println("lastName:", person1.lastName)
	fmt.Println("age:", person1.age)

}
func newPerson(firstName string, lastName string, age int) *person {
	var personTest = &person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}

	return personTest
}

type person struct {
	firstName string
	lastName  string
	age       int
}
