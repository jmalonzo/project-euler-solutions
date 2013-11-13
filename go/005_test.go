package main

import (
	"math"
	"testing"
)

const MaxInt = int(^uint(0) >> 1)

func smallestMultiple() int {
	divisible := false
	answer := 0
	for i := 20; i < MaxInt; i++ {
		for j := 1; j <= 20; j++ {
			divisible = math.Remainder(float64(i), float64(j)) == float64(0)
			if divisible == false {
				break
			}
		}
		if divisible == true {
			answer = i
			break
		}
	}
	return answer
}

func TestSmallestMultiple(t *testing.T) {
	x := smallestMultiple()
	answer := 232792560
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
}

