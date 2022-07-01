package main

import (
	"fmt"
)

func main() {
	var i interface{} //declaring empty interface
	describe(i)

	//We can provide any values to the empty interface
	i = 24
	describe(i) //giving int value to interface

	i = "hello"
	describe(i) //giving string value to interface
}

func describe(i interface{}) {
	fmt.Printf("Type:%T,Value :%v\n", i, i)
}
