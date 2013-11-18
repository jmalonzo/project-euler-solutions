package main

import (
	"testing"
)

func sumOfPrimesBelow(x int) int {
	sieve := make([]bool, x)
	for i:=0; i<x; i++ {
		sieve[i] = true
	}
	sieve[0] = false
	sieve[1] = false

	for i:=2; i<x; i++ {
		if sieve[i] == false {
			continue
		}
		for j:=2*i; j<x; j+=i {
			sieve[j] = false
		}
		
	}

	sum := 0
	for i:=2; i<x; i++ {
		if sieve[i] {
			sum += i;
		}
	}
	return sum
}

func TestSumOfPrimesBelow2M(t *testing.T) {
	x := sumOfPrimesBelow(2000000)
	answer := 142913828922
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
	
}
