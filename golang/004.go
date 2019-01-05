package main

import (
	"fmt"
	"strconv"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	max := 0
	for i := 999; i > 0; i-- {
		for j := 999; j > 0; j-- {
			cur := i * j
			prod := strconv.Itoa(cur)
			if prod == Reverse(prod) {
				if max < cur {
					max = cur
				}
			}
		}
	}
	fmt.Println(max)
}
