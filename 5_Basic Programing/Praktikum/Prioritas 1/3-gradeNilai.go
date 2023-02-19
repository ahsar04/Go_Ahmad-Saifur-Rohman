package main

import (
	"fmt"
)

func main() {
	nilai:=0
	fmt.Print("Masukan nilai :")
	fmt.Scanf("%d",&nilai)
	if nilai >=80 && nilai<=100{
		fmt.Print("Grade = A")
	}else if nilai >=65 && nilai<=79{
		fmt.Print("Grade = B")
	}else if nilai >=50 && nilai<=64{
		fmt.Print("Grade = C")
	}else if nilai >=35 && nilai<=49{
		fmt.Print("Grade = D")
	}else if nilai >=0 && nilai<=34{
		fmt.Print("Grade = E")
	}else{
		fmt.Print("Nilai Invalid")
	}
}
