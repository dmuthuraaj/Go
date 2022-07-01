package main

import "fmt"

func fibonacci(cap int, c chan int) {
	x := 0
	y := 1
	for i := 0; i < cap; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	channel := make(chan int, 20)
	go fibonacci(cap(channel), channel)
	for num := range channel {
		fmt.Println(num)
	}
}
