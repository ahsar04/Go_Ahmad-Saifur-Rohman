package main

import (
	"fmt"
)
func countSubarrays(numbers []int32, k int32) int64 {
    count := 0
    for i := 0; i < len(numbers); i++ {
        hasil := int32(1)
        for j := i; j < len(numbers); j++ {
            hasil *= numbers[j]
            
            if hasil <= k {
                count++
            } else {
                break
            }
        }
    }
    
    return int64(count)

}

func main() {
    fmt.Println(countSubarrays([]int32{1, 2, 3},4))
}
