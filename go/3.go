package main

import (
	"fmt"
	"math"
)

func main() {
	number := 600851475143
	lpf := 0
	nsqrt := int(math.Sqrt(float64(number)))
	sieve := make([]bool, nsqrt+1)

	// Initialize the sieve
	for i := 2; i <= nsqrt; i++ {
		sieve[i] = true
	}

	// Mark composites
	for i := 2; i <= nsqrt; i++ {
		if sieve[i] == true {
			ip := int(math.Pow(float64(i), 2))
			for j, k := ip, 1; j <= nsqrt; j = ip + (k * i) {
				sieve[j] = false
				k += 1
			}
		}
	}

	// Find the largest prime factor
	for i := nsqrt; i > 1; i-- {
		if sieve[i] == true && number%i == 0 {
			lpf = i
			break
		}
	}
	fmt.Println("The largest prime factor of", number, "is", lpf)
}
