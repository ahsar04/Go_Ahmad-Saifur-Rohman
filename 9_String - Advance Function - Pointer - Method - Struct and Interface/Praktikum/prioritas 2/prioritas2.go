package main

import (
	"fmt"
)

func caesar(offset int, input string) string  {
	var result string
	// input = strings.ToLower(input)
    if offset!=26 {
		if offset>26 {
			offset %=26
		}
		for _, data := range input {
			if data == ' ' {
				result += " "
				continue
			}
			newdata := int(data) + offset
			if newdata > int('z') {
				newdata -= 26
			}
			result += string(newdata)
		}
	}else{
		result = input
	}
    return result
}
func main()  {
	fmt.Println(caesar(26,"abc"))
	fmt.Println(caesar(2,"alta"))
	fmt.Println(caesar(10,"alteraacademy"))
	fmt.Println(caesar(1,"abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000,"abcdefghijklmnopqrstuvwxyz"))
}