package main

import (
	"fmt"
)

func main() {
	angka:=0
	fmt.Print("Masukan angka :")
	fmt.Scanf("%d",&angka)
	for i := 1; i <= angka; i++ {
		if angka%i==0 {
			fmt.Println(i)
		}
	}
}
