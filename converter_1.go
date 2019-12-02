package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "kinder"
	fmt.Println(myStrToInt(a))
	fmt.Println(myStrToInt2(a))
}

func myStrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	x, err := strconv.Atoi(nonFractionalPart[0])
	if str == "" {
		return 1, errors.New("empty string")
	}
	if x == 0 || err != nil {
		return 2, errors.New("bad string")
	}
	return strconv.Atoi(nonFractionalPart[0])
}
