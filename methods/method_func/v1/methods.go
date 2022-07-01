package main

import (
	"fmt"
)

type square struct {
	side int
}

func (f *square) sqfn() int {
	return f.side * f.side
}

func main() {
	s := square{10}
	fmt.Println(s.sqfn())
}
