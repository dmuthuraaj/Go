package main

import "fmt"

func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		f := i
		i, j = j, i+j
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
