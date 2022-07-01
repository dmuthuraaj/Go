package main

import (
	"fmt"
)

func main() {
	var _int int
	var _float float32
	var _bool bool
	var _ptr *int
	var _slice []int
	var _arr = [3]int{}

	fmt.Println("int", _int)
	fmt.Println("float", _float)
	fmt.Println("bool", _bool)
	fmt.Println("ptr", _ptr)
	fmt.Println("slice", _slice)
	fmt.Println("array", _arr)
}
