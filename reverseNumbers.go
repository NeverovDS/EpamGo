package main

import (
	"fmt"
)

func max(sl []int64) []int64 {
	k := make([]int64,0)

	for i, j := int64(0), int64(len(sl)-1); i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
		k = append(sl)
	}
	return k
}
func main() {
	s := []int64{1, 2, 5, 15,325,21,523,14,16,24,0}
	d := max(s)
	fmt.Println(d)
}