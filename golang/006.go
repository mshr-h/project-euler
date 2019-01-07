package main

import (
	"fmt"
)

func main() {
	var sum_of_sq = 0
	var sq_of_sum = 0

	for i := 1; i <= 100; i++ {
		sum_of_sq += i * i
		sq_of_sum += i
	}

	sq_of_sum *= sq_of_sum

	fmt.Println(sq_of_sum - sum_of_sq)
}
