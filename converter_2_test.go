package main

import (
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
	expected, nil := 1, error(nil)
	actual, nil := myStrToInt2(testString)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, actual)
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, nil)
	}
}
func Test_myStrToInt2V3(t *testing.T) {
	testString := "jSBV"
	expected, nil := 2, error(nil)
	actual, nil := myStrToInt2(testString)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, actual)
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, nil)
	}
}
