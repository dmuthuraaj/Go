package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var person1 person
	person2 := person{"muthu", 23}

	person1.name = "dhana"
	person1.age = 23

	fmt.Println(person1)
	_print(person2)
}

func _print(pers person) {
	fmt.Printf("Person name :%s\n", pers.name)
	fmt.Printf("Age : %d\n", pers.age)
}
