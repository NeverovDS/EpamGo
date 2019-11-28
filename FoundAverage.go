package main

import (
	"fmt"
)

func average(arr [] int) float64 {
	sum := 0
	n := len(arr)
	if len(arr) == 0 {
		return 0
	}
	if len(arr) != 0 {

		for i := 0; i < len(arr); i++ {
			sum += arr[i]
		}
	}
	return float64(sum) / float64(n)
}

func main() {
	someArray := [] int{}
	a := average(someArray)
	fmt.Println(a)
}
