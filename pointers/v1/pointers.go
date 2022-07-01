package main

import (
	"fmt"
)

const MAX int = 4

func main() {
	i := 0
	num := 10
	var num_ptr *int
	num_ptr = &num

	var pptr **int
	pptr = &num_ptr
	fmt.Printf("Address of single ptr %x = double ptr value:%x\n ", &num_ptr, pptr)
	fmt.Printf("Address of value %x = Single ptr value:%x\n", &num, *pptr)
	fmt.Printf("Value: by double ptr:%d ,by single ptr:%d\n ", **pptr, *num_ptr)

	numarr := [MAX]int{1, 2, 3, 4}
	var arr_ptr [MAX]*int

	fmt.Printf("Address:%x value:%d\n", &num, num)
	fmt.Printf("Address:%x value:%d\n", num_ptr, *num_ptr)

	arr_ptr[2] = &num

	for i = 0; i < MAX; i++ {
		arr_ptr[i] = &numarr[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Println(*arr_ptr[i])
	}
}
