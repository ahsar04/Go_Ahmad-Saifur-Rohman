package main

import (
	"fmt"
	"math"
)
func main() {
   fmt.Println(primeNumber(1500450271)) // true
   fmt.Println(primeNumber(1000000007)) // true
   fmt.Println(primeNumber(13))         // true
   fmt.Println(primeNumber(17))         // true
   fmt.Println(primeNumber(20))         // false
   fmt.Println(primeNumber(35))         // false
}
func primeNumber(number int) bool {
	
	if number < 2 {
		return false
	}else if number==2 {
		return true
	}
	quadrat := int(math.Sqrt(float64(number)))
	// fmt.Println(quadrat)
	for i:=2; i<=quadrat; i++{
		if number % i == 0 {
			return false
		}
	}
	return true
}