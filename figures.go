package main

import (
	"fmt"
)

const pi = 3.1415

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	a float64
}
type Circle struct {
	r float64
}

func (b Square) perimeter() float64 {
	b.a = 4 * b.a
	return b.a
}
func (b Square) area() float64 {
	b.a = b.a * b.a
	return b.a
}
func (b Circle) area() float64 {
	b.r = b.r * b.r * pi
	return b.r
}
func (b Circle) perimeter() float64 {
	b.r = b.r * 2 * pi
	return b.r
}

func main() {
	var s Figure = Square{5}
	var c Figure = Circle{3}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

}
