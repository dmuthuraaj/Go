package main

import (
	"fmt"
)

func main() {
	var luck_check map[int]string
	luck_check = make(map[int]string)
	var num int
	cars := make(map[int]string)

	luck_check[0] = "bad"
	luck_check[1] = "Not bad"
	luck_check[2] = "Good"
	luck_check[3] = "Better"

	cars[0] = "Audi"
	cars[1] = "BMW"
	cars[2] = "Ferrari"
	cars[3] = "Lamborghini"

	for k, v := range cars {
		fmt.Printf("value:%v car:%v\n", k, v)
	}

	fmt.Println("--------------luck_checker--------------")
	fmt.Println("press 9 to exit")
	for {
		fmt.Println("Enter no :")
		fmt.Scanf("%d", &num)
		_, ok := luck_check[num]
		if ok {
			fmt.Println(luck_check[num])
		} else if num == 9 {
			break
		} else {
			fmt.Println("Invalid value")
		}
	}
}
