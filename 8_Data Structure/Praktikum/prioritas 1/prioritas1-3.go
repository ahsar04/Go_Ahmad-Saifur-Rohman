package main

import (
	"fmt"
	"strconv"
)
func main()  {
	fmt.Println(munculSekali([]string{"1234123"}))
	fmt.Println(munculSekali([]string{"76523752"}))
	fmt.Println(munculSekali([]string{"12345"}))
	fmt.Println(munculSekali([]string{"1122334455"}))
	fmt.Println(munculSekali([]string{"0872504"}))
}
  

func munculSekali(angka []string)[]int  {
	counts := make(map[rune]int)
	result :=[]int{}
	for _, char := range angka[0] {
		counts[char]++
	}
	for char, count := range counts {
		if count <= 1 {
			// fmt.Println(string(char)) 
			strChar:= string(char)
			parseInt, err := strconv.Atoi(strChar)
			result = append(result, parseInt)

			if err != nil {
				fmt.Println("Data harus angka!")
				break
			}
		}
	}
	return result
}
