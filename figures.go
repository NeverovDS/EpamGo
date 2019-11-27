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
	perim := 4 * square.side
	return perim
}
func (square Square) area() float64 {
	ar := square.side * square.side
	return ar
}
func (circle Circle) area() float64 {
	ar := circle.r * circle.r * math.Pi
	return ar
}
func (circle Circle) perimeter() float64 {
	perim := circle.r * 2 * math.Pi
	return perim
}

func main() {
	var s Figure = Square{5}
	var c Figure = Circle{3}
	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

}
