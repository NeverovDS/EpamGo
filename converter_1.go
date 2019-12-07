package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "kinder"
	fmt.Println(myStrToInt(a))
	fmt.Println(myStrToInt2(a))
}

func myStrToInt(str string) (int, error) {
	if str == "" {
		return 1, nil
	}
	return strconv.Atoi(str)
}
