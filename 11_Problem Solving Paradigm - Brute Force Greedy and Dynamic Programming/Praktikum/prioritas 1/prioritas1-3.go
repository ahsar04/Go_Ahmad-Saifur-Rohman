package main

import (
	"fmt"
)

func fibo(number int)int  {
	if number == 0 || number==1{
        return number
    } else {
        return fibo(number-1) + fibo(number-2)
    }
}

func main()  {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}