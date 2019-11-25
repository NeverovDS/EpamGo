package main

import (
	"fmt"
	"math"
)


type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	side float64
}
type Circle struct {
	r float64
}

func (square Square) perimeter() float64 {
	square.side = 4 * square.side
	return square.side
}
func (square Square) area() float64 {
	square.side = square.side * square.side
	return square.side
}
func (circle Circle) area() float64 {
	circle.r = circle.r * circle.r * math.Pi
	return circle.r
}
func (circle Circle) perimeter() float64 {
	circle.r = circle.r * 2 * math.Pi
	return circle.r
}

func main() {
	var s Figure = Square{5}
	var c Figure = Circle{3}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

}
