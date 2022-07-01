package main

import "fmt"

func send(ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- "Hi"
	}
	close(ch)
}

func main() {
	my_channel := make(chan string)
	_chan := make(chan string)

	go send(my_channel)

	for {
		res, ok := <-my_channel
		if ok == false {
			fmt.Println("channel closed")
			break
		}
		fmt.Println(res, ok)
	}

	go func() {
		_chan <- "Hi"
		_chan <- "Hello"
		_chan <- "Welcome"
		close(_chan)
	}()

	for res := range _chan {
		fmt.Println(res)
	}
}
