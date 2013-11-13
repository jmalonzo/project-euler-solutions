package main

import "testing"

const Default = 1000

func Multiple(limit int) (result int) {
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	return result
}

func TestMultiple(t *testing.T) {
	x := Multiple(Default)
	answer := 233168
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
}
