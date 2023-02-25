package main

import (
	"fmt"
)
func main()  {
	fmt.Println(mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(mapping([]string{}))
}

func mapping(slice []string)map[string]int  {
	countWord:=make(map[string]int)
 	for _, v := range slice {
 		_, check := countWord[v]
		// fmt.Println(countWord)
 		if check {
 			countWord[v] += 1
 		} else {
 			countWord[v] = 1
 		}
 	}
 	return countWord
}
