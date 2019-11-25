package main

import (
	"fmt"
)

func reverseSlice(slice []int64) []int64 {
	result := make([]int64, 0)

	for i, j := int64(0), int64(len(slice)-1); i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
		result = append(slice)
	}
	return result
}
func main() {
	input := []int64{1, 2, 5, 15, 325, 21, 523, 14, 16, 24, 0}
	output := reverseSlice(input)
	fmt.Println(output)
}
