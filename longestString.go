package main

import (
	"fmt"
	"unicode/utf8"
)

func max(slice []string) string {
	count := 0
	longestWord := ""
	for i := 0; i < len(slice); i++ {

		if count < utf8.RuneCountInString(slice[i]) {
			count = utf8.RuneCountInString(slice[i])
			longestWord = slice[i]
		}
	}
	return longestWord

}

func main() {
	SliceString := []string{"one", "two"}
	a := max(SliceString)
	fmt.Println(a)
}
