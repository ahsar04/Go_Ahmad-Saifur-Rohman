package main

import (
	"fmt"
)



func getMinMax(numbers...*int)(min int, max int)  {
	var getMax, getMin int
    for i, num := range numbers {
        if i == 0 || *num > getMax {
            getMax = *num
        }
        if i == 0 || *num < getMin {
            getMin = *num
        }
    }
	// return 0,0
    return getMin, getMax
}

func main()  {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6 )
	fmt.Println("Output: ")
	fmt.Println("Nilai Min ", min)
	fmt.Println("Nilai Max ", max)
}