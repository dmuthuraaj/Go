package main

import (
	"fmt"
	"math"
)

func main() {
	var i, sum int

	for i = 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("square root of 4", math.Sqrt(4))
	fmt.Println("sum:", sum)

	/*
		for {                      //infinite loop
			fmt.Println(1)
		}
	*/

	if i < 10 {
		fmt.Println(i)
	} else {
		fmt.Println("High value")
	}

	i = 0
	for {
		fmt.Println(i)
		if i == 20 {
			break
		}
		i++
	}
}
