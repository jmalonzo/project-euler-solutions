package main

import "fmt"

func fibo(terms []int, limit int) []int {
	length := len(terms)
	r := terms[length-1] + terms[length-2]
	if r > limit {
		return terms
	}
	return fibo(append(terms, r), limit)
}

func main() {
	sum := 0
	for _, value := range fibo([]int{1, 2}, 4000000) {
		if value%2 == 0 {
			sum += value
		}
	}
	fmt.Println("The sum of the even terms of a fibo sequence under 4M is", sum)
}
