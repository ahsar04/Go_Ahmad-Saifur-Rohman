package main

import (
	"fmt"
	"math"
)

func simpleEquations(a, b, c int){
    x, y, z := 0, 0, 0
	found:=false
    for i := 1; i <= int(math.Sqrt(float64(b))); i++ {
        if b % i == 0 {
            j := b / i
            for k := -int(math.Sqrt(float64(c))); k <= int(math.Sqrt(float64(c))); k++ {
                if i != j && i != k && j != k && i + j + k == a && i * j * k == b && i*i + j*j + k*k == c {
                    x = i
                    y = j
                    z = k
					found = true
                }
            }
        }
    }
	if found==true {
		fmt.Println(x,y,z)
	}else{
		fmt.Println("no solution")
	}
}

func main() {
   simpleEquations(1,2,3)
   simpleEquations(6,6,14)
}
