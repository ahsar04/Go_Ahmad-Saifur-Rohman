package main

import (
	"fmt"
	"sync"
)

func countLetters(text string, wg *sync.WaitGroup, ch chan map[rune]int) {
	defer wg.Done()
	frequencies := make(map[rune]int)
	for _, letter := range text {
		frequencies[letter]++
	}
	ch <- frequencies
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse potenti. Nulla facilisi. Aliquam sodales ultricies dui."
	workers := 4
	ch := make(chan map[rune]int, workers)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		start := i * len(text) / workers
		end := (i + 1) * len(text) / workers
		go countLetters(text[start:end], &wg, ch)
	}

	wg.Wait()
	close(ch)

	frequencies := make(map[rune]int)
	for freqMap := range ch {
		for letter, freq := range freqMap {
			frequencies[letter] += freq
		}
	}

	fmt.Println("Frequency:")
	for letter, freq := range frequencies {
		fmt.Printf("%c: %d\n", letter, freq)
	}
}
