package main

import (
	"fmt"
)
func main() {
   fmt.Println(pow(2, 3))  // 8
   fmt.Println(pow(5, 3))  // 125
   fmt.Println(pow(10, 2)) // 100
   fmt.Println(pow(2, 5))  // 32
   fmt.Println(pow(7, 3))  // 343
}
func pow(x, n int) int {
	result := 1
    for i:=n; i>0; i--{
        if n % 2 == 0{
            x *= x
            n /= 2
		}else{
            result = result * x
            n -= 1
		}
	}
    return result
}

