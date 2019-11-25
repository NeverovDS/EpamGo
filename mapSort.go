package main

import (
	"fmt"
	"sort"
)

func printSorted(m1 map[int]string) {

	var keys []int
	for k := range m1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(m1[k])
	}
}

func main() {

	m := map[int]string{10: "a", 0: "b", 250: "c", 500: "d"}
	printSorted(m)

}
