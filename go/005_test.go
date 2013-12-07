package main

import (
	"math"
	"testing"
)

// From mathblog.dk/project-euler-problem-5
func generatePrimes(limit int) []int {
	var primes []int
	var j int
	isPrime := false

	primes = append(primes, 2)

	for i := 3; i <= limit; i += 2 {
		j = 0
		isPrime = true
		for primes[j]*primes[j] <= i {
			if i%primes[j] == 0 {
				isPrime = false
				break
			}
			j++
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func smallestMultiple() int {
	divisorMax := 20
	var primes []int = generatePrimes(divisorMax)
	result := 1

	for i := 0; i < len(primes); i++ {
		a := math.Floor(math.Log(float64(divisorMax)) / math.Log(float64(primes[i])))
		result = result * int(math.Pow(float64(primes[i]), a))
	}
	return result
}

func TestSmallestMultiple(t *testing.T) {
	x := smallestMultiple()
	answer := 232792560
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
}
