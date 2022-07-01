package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	go hello("Dhana", channel)
	for msg := range channel {
		fmt.Println(msg)
	}
}

func hello(name string, c chan string) {
	for i := 0; i < 4; i += 1 {
		msg := fmt.Sprintf("Hello%s", name)
		c <- msg
	}
	//closing channel otherwise
	//it runs in backround and
	//it occupies garbage storage
	close(c)
}
