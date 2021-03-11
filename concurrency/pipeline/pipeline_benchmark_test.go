package main

import "testing"

var min, max = 1, 100000

func BenchmarkSquareForArrPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		squareForArrPipeline(min, max)
	}
}

func BenchmarkSquareForArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		squareForArray(min, max)
	}
}
