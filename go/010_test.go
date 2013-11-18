package main

import (
	"math/big"
	"testing"
)

func sumOfPrimesBelow(x int64) int64 {
	var i int64
	var sum int64 = 0
	for i = 2; i < x; i++ {
		if i > 2 && i%2 == 0 {
			continue
		}

		j := big.NewInt(i)
		if !j.ProbablyPrime(10) {
			continue
		}
		sum += i
		
	}
	return sum
}

func TestSumOfPrimesBelow2M(t *testing.T) {
	x := sumOfPrimesBelow(2000000)
	var answer int64 = 142913828922
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
	
}
