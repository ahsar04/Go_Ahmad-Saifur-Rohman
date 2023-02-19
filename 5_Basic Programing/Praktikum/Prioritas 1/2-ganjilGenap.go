package main

import (
	"fmt"
)

func main() {
	nilai :=0
	fmt.Print("Masukan nilai :")
	fmt.Scanf("%d",&nilai)
	if nilai%2==0 {
		fmt.Print(nilai, " adalah :Genap")
	}else{
		fmt.Print(nilai, " adalah :Ganjil")
	}
}
