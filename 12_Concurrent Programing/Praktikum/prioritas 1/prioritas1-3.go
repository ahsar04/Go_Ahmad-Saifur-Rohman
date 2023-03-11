package main

import (
	"fmt"
	"time"
)

func multiples(c chan int) {
    i := 1
    for {
        c <- i * 3
        i++
        time.Sleep(1 * time.Second)
    }
}

func main() {
    channel := make(chan int, 5)
    go multiples(channel)

    for i := 0; i < 10; i++ {
        fmt.Println(<-channel)
    }
}
