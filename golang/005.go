package main

import (
	"fmt"
	"math"
)

var PARRAY = [8]int{2, 3, 5, 7, 11, 13, 17, 19}

func ToPrime(n int) map[int]int {
	var prime = make(map[int]int)
	for _, v := range PARRAY {
		prime[v] = 0
	}

	for _, v := range PARRAY {
		for n%v == 0 {
			n /= v
			prime[v] += 1
		}
	}

	return prime
}

func main() {
	var prime = make(map[int]int)
	for _, v := range PARRAY {
		prime[v] = 0
	}

	for x := 2; x <= 20; x++ {
		var a = ToPrime(x)
		for p, n := range a {
			if prime[p] < n {
				prime[p] = n
			}
		}
	}

	var prod = 1
	for i, v := range prime {
		prod *= int(math.Pow(float64(i), float64(v)))
	}
	fmt.Println(prod)
}
