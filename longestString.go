package main

import "fmt"

func max(slice []string) string {
	longestWord := 0
	b := ""
	for i := 0; i < len(slice); i++ {

		if longestWord < len(slice[i]) {
			longestWord = len(slice[i])
			b = slice[i]
		}
	}
	return b

}

func main() {
	SliceString := []string{"one", "two"}
	a := max(SliceString)
	fmt.Println(a)
}
