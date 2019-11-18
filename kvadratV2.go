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

func (NewPointX *Square) End() Point {
	x2 := NewPointX.start.x + NewPointX.a
	y2 := NewPointX.start.y - NewPointX.a
	return Point{x2, y2}
}
func (PerimeterDo *Square) Perimeter() int {
	return PerimeterDo.a * 4
}
func (AreaDo *Square) Area() int {
	return AreaDo.a * AreaDo.a
}

func main() {

	s := Square{start: Point{x: 2, y: 2}, a: 2}

	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())

}
