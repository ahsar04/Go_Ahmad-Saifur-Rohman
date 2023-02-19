package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	var isPalindrome bool
	// input data
    fmt.Print("Masukan data :")
	inputReader := bufio.NewReader(os.Stdin)
    inputData, _ := inputReader.ReadString('\n')

	// to lower case string
    inputData = strings.ToLower(inputData)

	// convert to rune
	runeData := []rune(inputData)
	// mapping kemudian dibalik per index
	for i, j := 0, len(runeData)-1; i < j; i, j = i+1, j-1 {

		runeData[i], runeData[j] = runeData[j], runeData[i]
	}
	
	fmt.Println("Captured :"+string(runeData))
	// maping rune(kurang begitu paham)
    inputData = strings.Map(func(r rune) rune {
        if r >= 'a' && r <= 'z' {
            return r
        }
        return -1
    }, inputData)
	
	// cek palindrome
    for i := 0; i < len(inputData)/2; i++ {
        if inputData[i] != inputData[len(inputData)-1-i] {
            isPalindrome=false
        }else{  
            isPalindrome=true
        }
    }
    if isPalindrome {
        fmt.Println("Palindrome.")
    } else {
        fmt.Println("Not Palindrome.")
    }
}