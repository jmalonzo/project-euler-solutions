package main

import (
	"testing"
)

func naivePythagoreanTriplet(limit int) (a,b,c int) {
	a, b, c = 0, 0, 0
	s := limit
	for a = 1; a < s/3; a++ {
		for b = a; b < s/2; b++ {
			c = s - a - b
			if (a*a) + (b*b) == (c*c) {
				return
			}
		}
	}
	return
}

func specialPythagoreanTriplet(limit int) (a, b, c int) {
	a, b, c = 0, 0, 0
	found := false
	for m:=1; m<limit; m++ {
		for n:=1; n<limit; n++ {
			if m < n || (m-n)%2 != 1 {
				continue
			}

			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n

			if a*a + b*b == c*c && a + b +c == limit {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return
}

func TestSpecialPythagoreanTriplet(t *testing.T) {
	a, b, c := specialPythagoreanTriplet(1000)
	x := a* b * c
	answer := 200 * 375 * 425
	if x != answer {
		t.Errorf("result: %v, want: %v", x, answer)
	}
}

