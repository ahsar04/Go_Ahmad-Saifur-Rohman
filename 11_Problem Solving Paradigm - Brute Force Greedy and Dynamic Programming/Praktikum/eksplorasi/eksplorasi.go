package main

import "fmt"

func intToRomawi(num int) string {
    romawi := []string{"I","IV","V","IX","X","XL","L","XC","C","CD","D","CM","M"}
    integer := []int{1,4,5,9,10,40,50,90,100,400,500,900,1000}
    result := ""
     for i := len(integer)-1; i >= 0; i-- {
        for num >= integer[i] {
            num -= integer[i]
            result += romawi[i]
        }
    }
    return result
}

func main() {
    fmt.Println(intToRomawi(4)) // Output: IV
    fmt.Println(intToRomawi(9)) // Output: IX
    fmt.Println(intToRomawi(8))
}
