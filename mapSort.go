package main

import (
	"fmt"
	"sort"
)

func printSorted(someMap map[int]string) {
	var keys []int
	for v := range someMap {
		keys = append(keys, v)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(someMap[v])
	}
}

func main() {
	input := map[int]string{10: "a", 0: "b", 250: "c", 500: "d"}
	printSorted(input)
}
