package main

import (
	"fmt"
	"strings"
)

func main() {
	_slice := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	_slice[0][0] = "X"
	_slice[1][0] = "X"
	_slice[2][0] = "O"
	_slice[0][1] = "O"
	_slice[2][2] = "X"

	for i := 0; i < len(_slice); i++ {
		fmt.Printf("%s\n", strings.Join(_slice[i], " "))
	}

	slice2 := make([]int, 3, 5)
	slice2[0] = 21
	slice2[1] = 2
	slice2[2] = 3
	slice2 = append(slice2, 4, 5, 6, 7)
	fmt.Println("slice2:", slice2, "length:", len(slice2), "capacity:", cap(slice2))

}
