package main

import (
	"fmt"
)

func _arr() {
	var num = []int{0, 1, 2, 3}
	fmt.Println(num)
}

//passing array as argument
func _rev_Arr(arr [2]string, size int) {
	var i int
	for i = size - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}

func main() {
	var i, length int
	print("array\n")
	cars := [2]string{"BMW", "AUDI"}
	fmt.Println(cars)
	//cars[2] = "FERRARI"        //we can't do this when we declared array without size
	//fmt.Println(cars)
	_arr()
	fmt.Printf("length of cars :%v \n", len(cars))
	length = len(cars)
	for i = 0; i < length; i++ {
		fmt.Printf("%v %v\n", i, cars[i])
	}
	//reversing array
	fmt.Print("Reversed String:\t")
	_rev_Arr(cars, length)
}
