package main

import (
	"fmt"
	"sync"
)

var (
	shared int  // variabel shared yang akan diakses oleh goroutine
	mu     sync.Mutex // objek mutex
)

func increment(wg *sync.WaitGroup, x int) {
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 5; i++ {
		shared += x
		fmt.Print(" ", shared)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg, 3)
	}

	wg.Wait() // menunggu semua goroutine selesai dieksekusi
}
