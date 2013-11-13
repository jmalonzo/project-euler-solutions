package main

import (
	"math/big"
	"testing"
)

func isPrime(x *big.Int) bool {
	return x.ProbablyPrime(10) == true
}

func nthPrime() int64 {
	p := big.NewInt(0)
	current := 0
	for current < 10001 {
		p = p.Add(p, big.NewInt(1))
		if isPrime(p) == true {
			current += 1
		}
	}
	return p.Int64()
}

func TestNthPrime(t *testing.T) {
	x := nthPrime()
	answer := int64(104743)
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
}
