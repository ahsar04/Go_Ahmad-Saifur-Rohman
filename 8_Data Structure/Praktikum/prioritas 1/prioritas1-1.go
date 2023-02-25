package main

import "fmt"
func main()  {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string {"eddie","steve","geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string {"jin","steve","bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string {"devil jin","yoshimitsu","alisa","law"}))
	fmt.Println(ArrayMerge([]string{}, []string {"devil jin","sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string {}))
	fmt.Println(ArrayMerge([]string{}, []string {}))

}

func ArrayMerge(arrayA, arrayB []string)[]string  {
	arrResult := make([]string,0)
	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(arrayB); j++ {
			if arrayA[i]==arrayB[j] {
				arrayB=append(arrayB[:j], arrayB[j+1:]...)
			}
		}
	}
	arrResult=append(arrayA,arrayB...)
	return arrResult
}
