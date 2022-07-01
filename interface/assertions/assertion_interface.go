//Like in map we are checking for is key is present or not
//here we are checking if the interface is provided with which type

package main

import "fmt"

func main() {

	var i interface{} = "Hello" //Declaring a interace and assining a string value

	//implementing interface as the string value given
	s := i.(string) //this statement asserts the value to be string and the type string is given to the variable s
	fmt.Println(s)

	s, ok := i.(string) //this statement aserts the value of type to interface and also check it is ok or not
	fmt.Println(s, ok)  //if the assertion is true then s returns the value of interface and the ok gives true

	//giving float value to interface but we declared as string interface
	f, ok := i.(float64) //for that we have to use new variable new variable for different assertions
	fmt.Println(f, ok)

	//triggering panic
	f = i.(float64) //this give a runtime error because we are giving float64 value but it has string
	fmt.Println(f)

}
