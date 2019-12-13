package main

import (
	"errors"
	"fmt"
)

func myStrToInt2(str string) (int, error) {
	var b int
	x, err := fmt.Sscanf(str, "%d", &b)
	if str == "" {
		return 0, errors.New("empty line")
	}
	if x == 0 || err != nil {
		return 0, errors.New("it is not an int value")
	}
	return b, nil
}
