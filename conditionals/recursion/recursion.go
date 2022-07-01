package main

import (
	"fmt"
)

func reverse_order(x int) int {
	if x >= 0 {
		fmt.Println(x)
	} else {
		return 0
	}
	return reverse_order(x - 1)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println("Reverse order")
	reverse_order(5)
	fmt.Println("Recursion")
	fmt.Println(factorial_recursion(4))
}
