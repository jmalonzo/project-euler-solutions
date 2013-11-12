package main

import (
	"fmt"
	"math/big"
)

func isPrime(x *big.Int) bool {
	return x.ProbablyPrime(10) == true
}

func main() {
	p := big.NewInt(0)
	current := 0
	for current < 10001 {
		p = p.Add(p, big.NewInt(1))
		if isPrime(p) == true {
			current += 1
		}
	}
	fmt.Println("The 10 001st prime number is", p)
}
