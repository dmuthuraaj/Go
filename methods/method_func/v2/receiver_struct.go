package main

import "log"

type person struct {
	firstname string
	lastname  string
}

func (p *person) printfirstname() string {
	return p.firstname
}

func main() {
	var person1 person
	person1.firstname = "dhana"

	person2 := person{
		firstname: "muthu",
		lastname:  "raaj",
	}

	log.Println("person1", person1.printfirstname())
	log.Println("person2", person2.printfirstname())
}
