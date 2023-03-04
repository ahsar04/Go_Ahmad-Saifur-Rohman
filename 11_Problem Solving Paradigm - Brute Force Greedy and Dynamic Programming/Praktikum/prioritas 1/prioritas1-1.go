package main

import (
	"fmt"
)

func getBiner(n int) []int{
    ans := make([]int, n+1)
    for i := 0; i <= n; i++ {
        biner := 0
        base := 1
        num := i
        for num > 0 {
            remainder := num % 2
            biner += remainder * base
            num /= 2
            base *= 10
        }
        ans[i] = biner
    }
    return ans
}

func main()  {
	fmt.Println(getBiner(2))
	fmt.Println(getBiner(3))
}