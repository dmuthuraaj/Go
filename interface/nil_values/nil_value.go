package main

import (
	"fmt"
)

//structure for interface
type Person struct {
	name string
}

//declaring interface
type Print interface {
	print()
}

//defining method of interface
func (p *Person) print() {
	if p == nil { //checking the value is nil or not
		fmt.Println(nil) //if nil it should print this
		return
	}
	fmt.Println(p.name)
}

func main() {
	//declaring interface
	var p Print
	//Giving nil value to interface
	var per *Person //intitialising nil struct for passing to interface
	p = per         //for pointer receiver we have to give address
	describe(p)     //for describing interface type and value
	p.print()       //calling interface method

}

//printing type and value of interface
func describe(i Print) {
	fmt.Printf("Type:%T Value:%v\n", i, i)
}
