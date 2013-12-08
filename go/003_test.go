package main

import (
	"math"
	"testing"
)

func largestPrimeFactor(number int) int {
	lpf := 0
	nsqrt := int(math.Sqrt(float64(number)))
	sieve := make([]bool, nsqrt+1)

	// Initialize the sieve
	for i := 2; i <= nsqrt; i++ {
		sieve[i] = true
	}

	// Mark composites
	for i := 2; i <= nsqrt; i++ {
		if sieve[i] == false {
			continue
		}
		for j := 2 * i; j < nsqrt; j += i {
			sieve[j] = false
		}
	}

	// Find the largest prime factor
	for i := nsqrt; i > 1; i-- {
		if sieve[i] == true && number%i == 0 {
			lpf = i
			break
		}
	}

	return lpf
}

func TestGetLargestPrimeFactor(t *testing.T) {
	x := largestPrimeFactor(600851475143)
	answer := 6857
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
}
