package main

import (
	"fmt"
)

func main() {
// L = 1/2 × t ( a + b )
/*===mau pakai scanner tapi masih ada kendala yaitu apabila dijalankan scanner hanya muncul 2X===*/
	// var L, t,a,b int
	// fmt.Print("Masukan tinggi :")
	// fmt.Scanf("%d" ,&t)

	// fmt.Print("Masukan a :")
	// fmt.Scanf("%d",&a)
	// scanner yang dibawah ini tidak muncul
	// fmt.Print("Masukan b :")
	// fmt.Scanf("%d",&b)
	// L = t*(a+b)*1/2
	// fmt.Println("Luas segitiganya adalah : ",L)

	// alhasil nilainya input manual
	var L int
	t:=12
	a:=8
	b:=16
	L = t*(a+b)*1/2
	fmt.Println("Luas trapesium adalah : ",L)
}
