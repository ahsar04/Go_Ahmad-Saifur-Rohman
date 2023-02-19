package main

import (
	"fmt"
)

func main() {
	nilai:=0
	data:=0
	data2:=0
	fmt.Print("Masukan nilai :")
	fmt.Scanf("%d",&nilai)
	data2=nilai-1
	for i := 0; i < nilai; i++ {
		// fmt.Print(i)
		for j := 0; j < data2; j++ {
			fmt.Print(" ")
			
		}
		data2-=1
		for j := 0; j <= data; j++ {
			if j%2==0 {
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		data+=2
		fmt.Println()
	}
}
