package main

import "fmt"

func factorial(i uint) uint {
	var r uint = 1
	var k uint
	if i == 0 {
		return 1
	}

	for k = 2; k <= i; k++ {
		r *= k

	}
	return r
}

func main() {
	fmt.Println(factorial(33))
}
