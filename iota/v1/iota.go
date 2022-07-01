package main

import (
	"fmt"
)

const (
	Monday int = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println("Hi", Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
}
