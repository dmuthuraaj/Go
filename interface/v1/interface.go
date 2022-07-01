package main

import "fmt"

type _shape interface {
	area() float64
}

type _rectangle struct {
	length, breadth float64
}

type _square struct {
	side float64
}

func (square _square) area() float64 {
	return square.side * square.side
}

func (rect _rectangle) area() float64 {
	return rect.breadth * rect.length
}

func getArea(shape _shape) float64 {
	return shape.area()
}

func main() {

	square := _square{side: 10}
	fmt.Println("Area of square:", getArea(square))

	rectangle := _rectangle{breadth: 2, length: 4}
	fmt.Println("Area of rectangle:", getArea(rectangle))
}
