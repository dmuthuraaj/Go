package main

import (
	"fmt"
)

//for reversing the given string
func Reverse_Runes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < (len(r) / 2); i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	var str string = "hi"
	str = "hello"
	fmt.Println(str)
	fmt.Println(Reverse_Runes("Hello,wolrd"))
}
