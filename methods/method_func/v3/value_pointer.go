package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

type Show interface {
	showFirstName() string
	showLastName() string
}

func (p Person) showFirstName() string {
	return p.firstName
}

func (p *Person) showLastName() string {
	return p.lastName
}

func main() {
	var s Show
	p := Person{"Muthu", "Raaj"}

	//we are Passing pointer
	//it is working for both methods
	s = &p
	fmt.Println(s.showFirstName())
	fmt.Println(s.showLastName())

	//we can't give value for interface s
	//because showLastName() has a pointer receiver
	//(s = p) is not accepted

	person := Person{
		"Dhana",
		"Muthu",
	}

	fmt.Println("First Name:", person.showFirstName())
	fmt.Println("second Name: ", person.showLastName())
}
