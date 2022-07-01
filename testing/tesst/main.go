package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sqrt(9)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func sqrt(value float64) (float64, error) {
	var result float64
	if value < 0 {
		return 0, errors.New("negative value")
	}
	result = math.Sqrt(value)
	return result, nil
}
