package main

import (
	"fmt"
)

func myStrToInt2(str string) (int, error) {
	var b int
	x, err := fmt.Sscanf(str, "%d", &b)
	if str == "" {
		return 1, error(nil)
	}
	if x == 0 || err != nil {
		return 2, error(nil)
	}
	return b, nil
}
