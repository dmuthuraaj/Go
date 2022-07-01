package main

import (
	"fmt"
	"time"
)

func print(s string) {
	fmt.Printf("It's %v\n", s)
}

func main() {
	//getting the full time with day year time all
	t := time.Now()
	//getting today
	day := t.Weekday()

	fmt.Println("time:", t)
	fmt.Println("hour:", t.Hour())

	switch day {
	case time.Monday:
		print("monday")
	default:
		print(day.String())
	}

}
