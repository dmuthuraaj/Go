package main

import "fmt"

func deliver(fn func(int, int) int) int {
	return fn(4, 5)
}

func main() {
	func1 := func(x, y int) int {
		return x + y
	}

	fmt.Println(func1(1, 2))
	fmt.Println(deliver(func1))
}
