package main

import "testing"

func fibo(terms []int, limit int) []int {
	length := len(terms)
	r := terms[length-1] + terms[length-2]
	if r > limit {
		return terms
	}
	return fibo(append(terms, r), limit)
}

func runFibo() int {
	sum := 0
	for _, value := range fibo([]int{1, 2}, 4000000) {
		if value%2 == 0 {
			sum += value
		}
	}
	return sum
}

func TestFibo(t *testing.T) {
	x := runFibo()
	answer := 4613732
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
}
