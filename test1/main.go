package main

import (
	"fmt"
	"sync"
)

func generateNumbers(start, end int, oddch chan<- int, evench chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evench <- i
		} else {
			oddch <- i
		}
	}
	close(oddch)
	close(evench)
}

func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup

	// Create two channels for odd and even numbers
	oddCh := make(chan int)
	evenCh := make(chan int)

	// Number range
	start := 1
	end := 10

	// Use goroutines to generate odd and even numbers concurrently
	wg.Add(1)
	go generateNumbers(start, end, oddCh, evenCh, &wg)

	// Use a goroutine to print odd numbers
	wg.Add(1)
	go printNumbers(oddCh, &wg)

	// Use a goroutine to print even numbers
	wg.Add(1)
	go printNumbers(evenCh, &wg)

	wg.Wait()
}
