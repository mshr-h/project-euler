package main

import (
	"fmt"
	"math/big"
)

func main() {
	from := 100
	factorial := big.NewInt(int64(from))
	for i := 1; i < from; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}

	sum := 0
	for i := 0; i < len(factorial.String()); i++ {
		n := factorial.String()[i] - '0'
		sum += int(n)
	}

	fmt.Println(sum)
}
