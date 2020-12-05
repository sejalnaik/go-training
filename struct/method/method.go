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

	fmt.Println("person1:", person1)
	fmt.Println("adress of person1:", &person1)
	fmt.Println("firstName:", person1.firstName)
	fmt.Println("lastName:", person1.lastName)
	fmt.Println("age:", person1.age)
	fmt.Println("full name:", person1.getFullName())
}

//value reciever
/*func (p person) getFullName() string {
	return p.firstName + " " + p.lastName
}*/

//ponter receviver
func (p *person) getFullName() string {
	return p.firstName + " " + p.lastName
}

type person struct {
	firstName string
	lastName  string
	age       int
	addresses []string
}
