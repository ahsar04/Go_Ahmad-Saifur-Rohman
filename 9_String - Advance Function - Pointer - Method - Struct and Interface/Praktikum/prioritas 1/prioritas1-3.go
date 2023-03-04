package main

import (
	"fmt"
	"strings"
)

func compare(a, b string)string  {
	a=strings.ToUpper(a)
	b=strings.ToUpper(b)
	var isExists1 = strings.Contains(a, b)
	var isExists2 = strings.Contains(b, a)
    if isExists1 {
		return b
	}else if isExists2 {
		return a
	}else {
		return "Invalid input"
	}
}

func main()  {
	fmt.Println(compare("AKA", "AKASHI"))
	fmt.Println(compare("KANGORO", "KANG"))
	fmt.Println(compare("KI", "KIJANG"))
	fmt.Println(compare("KUPU-KUPU", "KUPU"))
	fmt.Println(compare("ILALANG", "ILA"))
	fmt.Println(compare("Saifur", "Sai"))
	fmt.Println(compare("SAIFUR", "Sai"))

}