package main

import (
	"flag"
	"fmt"
)

func IsPrime(n int) bool {
	if n%2 == 0 {
		return false
	}

	for i := 3; i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Usage
	// go run 007.go -n 10001
	var np = flag.Int("n", 6, "nth prime number")
	flag.Parse()

	var cnt = 0

	for i := 0; ; i++ {
		if IsPrime(i) {
			cnt++
		}
		if cnt == *np {
			fmt.Println(*np, "th prime number is ", i)
			break
		}
	}
}
