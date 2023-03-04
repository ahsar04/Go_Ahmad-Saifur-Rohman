package main

import (
	"fmt"
	// "math"
)
func simpleEquations(a,b,c int)  {
	// for x := int(math.Sqrt(float64(c))); x <= a/3; x++ {
	// 	for y := x; y <= (a-x)/2; y++ {
	// 		z := a - x - y
	// 		if x*y*z == b && x*x+y*y+z*z == c {
	// 			fmt.Print(x, y, z)
	// 		}
	// 	}
	// }
	fmt.Println("no solution")
}
func main() {
	simpleEquations(1, 2, 3)
	simpleEquations(6, 6, 14)
}
