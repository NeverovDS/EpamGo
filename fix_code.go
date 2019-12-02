package main

import (
	"encoding/json"
	"errors"
)

func multiplyByTwo(k int) error {
	k *= 2
	return nil
}

func printMoreTen(k int64) error {
	if k < 10 {
		return errors.New("k must be > 10")
	}
	return nil
}

func dejson() error {
	type jsStruct struct {
		Data int  `json:"data"`
		OK   bool `json:"ok"`
	}
	in := []byte(`{"data": 13, "ok": true}`)
	var out jsStruct
	if err := json.Unmarshal(in, &out); err != nil {
		panic(err)
	}
	return nil
}

func main() {
	var r int64 = 11
	var k = 10
	err := multiplyByTwo(int(r))
	err2 := multiplyByTwo(k)
	if err = printMoreTen(r); err != nil {
		panic(err)
	}
	if err2 = printMoreTen(r); err2 != nil {
		panic(err)
	}
}
