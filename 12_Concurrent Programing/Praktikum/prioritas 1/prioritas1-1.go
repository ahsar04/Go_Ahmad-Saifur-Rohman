package main

import (
    "fmt"
    "time"
)

func multiples(number int) {
    i := 1
    for {
        fmt.Printf("%d ", i*number)
        i++
        time.Sleep(3 * time.Second)
    }
}

func main() {
    go multiples(4)
    var input string
    fmt.Scanln(&input)
}
