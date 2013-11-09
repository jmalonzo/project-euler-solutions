
package main

import (
	"fmt"
	"math"
)

const MaxInt = int(^uint(0) >> 1) 

func main() {
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
	fmt.Println("Smallest number divisible by numbers 1..20 is", answer)
}
