package main

import "fmt"


type point struct {
	 x, y int
}

type square struct {
	start point
	End   point
}


func (x1 point) prevrashenee() int , int {
	return (x1.x + 5) , (x1.y - 5)
}

func main() {
	x1 : = point{3,2}
	
	fmt.Println(x1.prevrashenee())

}
