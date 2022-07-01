package main

import (
	"fmt"
)

func _return(x int) int {
	return x
}

func swap(x, y string) (string, string) {
	return y, x
}

var glob float32 = 0.07 //declaration & initialisation from outside the function

func main() {
	var number int = 89 //dec & initialising with type and value
	var _bool bool = 8 > 10
	var person string //dec without initialisation

	value := 100     //dec & initialisation
	person = "Dhana" //initialising valus to variable

	var a, b, c, d int = 2, 4, 6, 8 //decalring multiple variable with type and value(one type only be used)
	var e, f = 10, "Hi"             //decalring without type
	g, h := 12, "Hello"

	var (
		i int
		j int    = 14
		k        = 16
		l string = "Welcome"
		m        = "You"
	)

	const PI float32 = 3.14 //typed const
	const NAME = "name"     //untyped constants

	//printing values
	fmt.Print(glob, "\n", number, "\n", person, " ")
	fmt.Println(_bool)
	fmt.Printf("%c\n", value) //giving int format specifier to character value
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m)
	fmt.Println(PI, NAME)
	fmt.Printf("Name:%v ,Type : %T\n", NAME, NAME)

	//returned value from function
	fmt.Printf("Returned value from function : %v\n", _return(55))

	//swapping using swap function
	n := "Hi"
	o := "Hello"
	fmt.Printf("values x :%s y:%s\n", n, o)
	p, q := swap(n, o)
	fmt.Printf("Swapped values x :%s y:%s\n", p, q)

	//getting input from user
	fmt.Println("Enter your name")
	var _name string
	fmt.Scanln(&_name)
	fmt.Printf("Hi, %s", _name)
}
