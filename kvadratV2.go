package main

import (
	"fmt"
)

type Square struct {
	start Point
	a     int
}

type Point struct {
	x, y int
}

func (X *Square) End() Point {
	x2 := X.start.x + X.a
	y2 := X.start.y - X.a
	return Point{x2, y2}
}
func (X *Square) Perimeter() int {
	return X.a * 4
}
func (X *Square) Area() int {
	return X.a * X.a
}

func main() {

	s := Square{start: Point{x: 2, y: 2}, a: 2}

	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())

}
