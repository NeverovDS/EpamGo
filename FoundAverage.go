package main

import (
	"fmt"
)

func average(arr [] int) float64 {
	sum := 0
	n := len(arr)
	if n == 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(n)
}

func main() {
	someArray := [] int{12, 2, 3}
	a := average(someArray)
	fmt.Println(a)
}
