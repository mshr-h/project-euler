package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(2)
	y := big.NewInt(1000)
	x.Exp(x, y, nil)

	sum := 0
	for i := 0; i < len(x.String()); i++ {
		n := x.String()[i] - '0'
		sum += int(n)
	}

	fmt.Println(sum)
}
