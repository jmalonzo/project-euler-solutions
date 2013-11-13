package main

import (
	"strconv"
	"testing"
)

func palindrome(num int) int {
	s := strconv.Itoa(num)
	l := len(s)
	for i := 0; i < l/2; i += 1 {
		if s[i] != s[l-i-1] {
			return 0
		}
	}
	return num
}

func largestPalindrome() int {
	start, end := 999, 99
	lp, num, p := 0, 0, 0
	for i := start; i > end; i = i - 1 {
		for j := start; j > end; j = j - 1 {
			num = i * j
			p = palindrome(num)
			if p > 0 && p > lp {
				lp = p
			}
		}
	}
	return lp
}

func TestLargestPalindrome(t *testing.T) {
	x := largestPalindrome()
	answer := 906609
	if x != answer {
		t.Errorf("result = %v, want %v", x, answer)
	}
}
