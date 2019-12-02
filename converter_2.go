package main

import (
	"errors"
	"fmt"
)

func myStrToInt2(str string) (int, error) {
	var b int
	x, err := fmt.Sscanf(str, "%d", &b)
	if str == "" {
		return 1, errors.New("empty string")
	}
	if x == 0 || err != nil {
		return 2, errors.New("bad string")
	}
	fmt.Sscanf(str, "%d", &b)
	return b, nil
}
