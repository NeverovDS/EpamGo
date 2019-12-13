package main

import (
	"errors"
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
		return 0, errors.New("empty line")
	}
	return strconv.Atoi(str)
}
