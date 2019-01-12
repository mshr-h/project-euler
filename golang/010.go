package main

import (
        "fmt"
)

func main() {
        var memo []int
        memo = append(memo, 2)
        var sum = int64(2)
        for i := 3; i < 2000000; i += 2 {
                var prime = true
                for _, v := range memo {
                        if i%v == 0 {
                                prime = false
                                break
                        }
                }
                if prime {
                        memo = append(memo, i)
                        sum += int64(i)
                }
        }
        fmt.Println(sum)
}
