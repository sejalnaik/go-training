package main

import "fmt"

func main() {
	var person1 = person{
		firstName: "sejal",
		lastName:  "naik",
		age:       26,
		addresses: []string{
			"mumbai",
			"delhi",
		},
	}

	var person2 = person{
		firstName: "rachel",
		lastName:  "green",
		age:       30,
		addresses: []string{
			"new york",
			"florida",
		},
	}

	var person3 = person{
		firstName: "ross",
		lastName:  "geller",
		age:       40,
		addresses: []string{
			"mumbai",
			"goa",
		},
	}

	people := []person{person1, person2, person3}

	for _, singlePerson := range people {
		fmt.Println("firstName:", singlePerson.firstName)
		fmt.Println("lastName:", singlePerson.lastName)
		fmt.Println("age:", singlePerson.age)

		for _, address := range singlePerson.addresses {
			fmt.Print("address:", address, " ")
		}
		fmt.Print("\n")
	}
}

type person struct {
	firstName string
	lastName  string
	age       int
	addresses []string
}
