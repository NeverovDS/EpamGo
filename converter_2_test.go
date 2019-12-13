package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_myStrToInt2(t *testing.T) {
	testData := "23"
	expected, nil := 23, error(nil)
	actual, nil := myStrToInt2(testData)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, actual)
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, nil)
	}
}
func Test_myStrToInt2V2(t *testing.T) {
	testString := ""
	expected, err := 0, errors.New("empty line")
	actual, err := myStrToInt2(testString)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, actual)
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, err)
	}
}
func Test_myStrToInt2V3(t *testing.T) {
	testString := "jSBV"
	expected, err := 0, errors.New("it is not an int value")
	actual, err := myStrToInt2(testString)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, actual)
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, err)
	}
}
