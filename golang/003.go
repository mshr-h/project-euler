package main

import (
	"fmt"
)

func main() {
	target := int64(600851475143)
	for i := int64(2); i < target; i++ {
		if target%i == 0 {
			fmt.Println(i)
			target /= i
		}
	}
	fmt.Println(target)
}
