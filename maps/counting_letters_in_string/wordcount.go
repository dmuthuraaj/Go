package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	_slice := strings.Fields(s)
	map1 := make(map[string]int)
	for _, val := range _slice {
		//map1[val]++
		count, ok := map1[val]
		if ok {
			map1[val] = count + 1
		} else {
			map1[val] = 1
		}
	}
	return map1
}

func main() {
	_map := WordCount("I am dhana I am intern")
	fmt.Println(_map)
}
