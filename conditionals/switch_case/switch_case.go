package main

import "fmt"

func main() {
	var i int = 2

	switch i {
	case 1:
		fmt.Println("i is set to 1")
	case 2:
		fmt.Println("i is set to 2")
	case 3:
		fmt.Println("i is set to 3")
	default:
		fmt.Println("i is set to something else")
	}

}
