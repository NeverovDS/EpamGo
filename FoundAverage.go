package main

import "fmt"

func average(arr [] int) float64 {
	 Sum  := 0
	N := len(arr)
	for i := 0; i < len(arr) ; i++ {
	Sum += arr[i]
	}
	return float64(Sum) / float64(N)
	}
	func main() {

	 SomeArray := [] int{4,2,3,1}
	 a := average(SomeArray)

		fmt.Println(a)
	}

