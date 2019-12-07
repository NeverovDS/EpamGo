package main

import (
	"testing"
)

func Benchmark_myStrToInt(b *testing.B) {
	testData := "12"
	for i := 0; i < b.N; i++ {
		myStrToInt(testData)
	}
}
func Benchmark_myStrToInt2(b *testing.B) {
	testData := "12"
	for i := 0; i < b.N; i++ {
		myStrToInt2(testData)
	}
}
