package personstruct

func NewPerson(firstName string, lastName string, age int) *person {
	var personTest = &person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return personTest
}

type person struct {
	FirstName string
	LastName  string
	Age       int
}
