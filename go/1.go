package main

import (
	"fmt"
)

const Default = 1000

func multiple(limit int) (result int) {
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	return result
}

func main() {
	r := multiple(Default)
	fmt.Println("The result is: ", r)
}
