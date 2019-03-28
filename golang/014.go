package main

import "fmt"

// collatz() takes starting number and returns number of collatz sequence chain
func collatz(start int) int {
	chain := 1
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = start*3 + 1
		}
		chain += 1
	}
	return chain
}

func main() {
	chain := 1
	start := 8

	maxChain := 1
	maxStart := 1

	for i := 1; i < 1000000; i++ {
		chain = collatz(i)
		start = i

		if chain > maxChain {
			maxChain = chain
			maxStart = start
		}
	}
	fmt.Println("maxChain:", maxChain, "maxStart:", maxStart)
}
