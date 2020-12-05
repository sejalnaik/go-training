package main

import "fmt"

func main() {
	var employee1 = employee{
		employeeID: 100,
		person: person{
			firstName: "sejal",
			lastName:  "naik",
			age:       26,
		},
	}

	fmt.Println("employee:", employee1)
	fmt.Println("adress of employee1:", &employee1)
	fmt.Println("ID of employee1:", employee1.employeeID)
	fmt.Println("firstName:", employee1.person.firstName)
	fmt.Println("lastName:", employee1.person.lastName)
	fmt.Println("age:", employee1.person.age)
}

type employee struct {
	employeeID int
	person
}

type person struct {
	firstName string
	lastName  string
	age       int
}
