package main

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func Test_myStrToInt(t *testing.T) {
	testData := "23.5"
	expected, nil := strconv.Atoi(testData)
	actual, nil := myStrToInt(testData)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, actual)
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, nil)
	}
}
func Test_myStrToIntV2(t *testing.T) {
	testString := ""
	expected, nil := 0, errors.New("empty line")
	actual, nil := myStrToInt(testString)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, actual)
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, nil)
	}
}
func Test_myStrToIntV3(t *testing.T) {
	testString := "jSBV"
	expected, nil := strconv.Atoi(testString)
	actual, nil := myStrToInt(testString)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, actual)
		t.Errorf("Not equal\nexpected - %v\nactual - %v", expected, nil)
	}
}
