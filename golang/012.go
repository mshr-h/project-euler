package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

func divisors(x int64) int64 {
	var num = int64(0)
	for i := int64(1); i <= int64(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			if x/i == i {
				num++
			} else {
				num += 2
			}
		}
	}
	return num
}

func main() {
	ch := make(chan int, runtime.NumCPU())
	triangleNumber := int64(0)
	for i := int64(1); ; i++ {
		ch <- 1
		triangleNumber += i
		go func(x int64) {
			if divisors(x) > int64(500) {
				fmt.Println(x)
				os.Exit(0)
			}
			<-ch
		}(triangleNumber)
	}
}
