package main

import (
	"fmt"
	"math"
)

type value float64

type person struct {
	name string
}

type Print interface {
	print()
}

func (f value) print() {
	fmt.Println(f)
}

func (s person) print() {
	fmt.Println(s.name)
}

func main() {
	var i Print

	//providing float value for interface
	var f value = math.Pi
	i = f
	desc(i)
	i.print()

	//providing struct for interface
	p := person{"Dhana"}
	i = p
	desc(i)
	i.print()
}

//Printing the value of interface and it's type
//we will get waht type we are giving like float and struct
//that we are passing here
func desc(i Print) {
	fmt.Printf("Type:%T,Value:%v\n", i, i)
}
