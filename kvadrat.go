package main

import (
	"fmt"
)

type Square struct {
	start    Point
	a        int
	End      PointEnd
	Perimeter int
	Area     int
}
type Point struct {
	x, y int
}
type PointEnd struct {
	x1, y1 int
}

func (NewPointX *Square) BuildEnd() PointEnd {
	x1:=  NewPointX.start.x + NewPointX.a
	y1:= NewPointX.start.y - NewPointX.a
	return PointEnd{x1,y1}
}

func (PerimetrGO *Square) BuildPerimeter() int {
	return 4 * PerimetrGO.a
}
func (AreaGO *Square) BuildArea() int {
	return AreaGO.a * AreaGO.a
}
func main() {

	s := Square{start: Point{x: 2, y: 2}, a: 2}
	s = Square{End: s.BuildEnd(), Perimeter: s.BuildPerimeter(), Area: s.BuildArea()}
	fmt.Println(s.End)
	fmt.Println(s.Perimeter)
	fmt.Println(s.Area)

}
