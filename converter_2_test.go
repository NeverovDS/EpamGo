package main
import (
	"errors"
	"reflect"
	"testing"
)

func Test_myStrToInt2(t *testing.T){
	testData := "23"
	expected , _ := 23 , error(nil)
	actual , _ := myStrToInt2(testData)
	if !reflect.DeepEqual(expected,actual){
		t.Errorf("Not equal\nexpected - %v\nactual - %v" , expected , actual )
	}
}
func Test_myStrToInt2V2(t *testing.T){
	testString := ""
	expected , _ := 1, errors.New("empty string")
	actual , _ := myStrToInt2(testString)
	if !reflect.DeepEqual(expected,actual){
		t.Errorf("Not equal\nexpected - %v\nactual - %v" , expected , actual )
	}
}
func Test_myStrToInt2V3(t *testing.T){
	testString := "jSBV"
	expected , _ :=  2, errors.New("bad string")
	actual , _ := myStrToInt2(testString)
	if !reflect.DeepEqual(expected,actual){
		t.Errorf("Not equal\nexpected - %v\nactual - %v" , expected , actual )
	}
}

