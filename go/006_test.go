package main

import (
	"math"
	"testing"
)

func sumOfTheSquares(x int64) int64 {
	var r int64 = 0
	var i int64
	for i = 1; i <= x; i++ {
		r += int64(math.Pow(float64(i), float64(2)))
	}
	return r
}

func squaresOfTheSum(x int64) int64 {
	var r int64 = 0
	var i int64
	for i = 1; i <= x; i++ {
		r += i
	}
	return int64(math.Pow(float64(r), float64(2)))
}

func sumSquareDifference() int64 {
	var (
		num    int64 = 100
		result int64 = squaresOfTheSum(num) - sumOfTheSquares(num)
	)
	return result
}

func TestSumSquareDifference(t *testing.T) {
	var answer int64 = 25164150
	x := sumSquareDifference()
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
}
