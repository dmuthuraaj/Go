package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func power(base, pow float64) float64 {
	return math.Pow(base, pow)
}

func main() {
	//Pi value
	fmt.Println("Pi value:", math.Pi)
	//printing square root of number
	fmt.Println("Square root of 36:", math.Sqrt(36))
	//using ramd.seed for changing the valu of rand.intn value everytime
	rand.Seed(int64(time.Now().Second()))
	//printing random value
	fmt.Println("Random number within 100:", rand.Intn(100))
	//printing the power
	fmt.Println("6 to the power 2:", power(6, 2))
}
