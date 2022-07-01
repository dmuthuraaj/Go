package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 20, 30, 40, 50, 60}
	printslice(nums)

	nums = append(nums, 70, 80)
	printslice(nums)

	copy_nums := make([]int, len(nums), (cap(nums))*2)
	copy(copy_nums, nums)
	printslice(copy_nums)
	fmt.Println(nums[:3])

}

func printslice(x []int) {
	fmt.Printf("length:%d,capacity:%d,slice:%v\n", cap(x), len(x), x)
}
