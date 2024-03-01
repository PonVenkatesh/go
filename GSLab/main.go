package main

import (
	"fmt"
	"sync"
)

func print_fruit_info(fruit_info map[string]string, f chan string, c chan string) {
	for fruit, color := range fruit_info {
		f <- fruit
		c <- color
	}
	close(f)
	close(c)
}

func printFruit(f <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for fruit := range f {
		fmt.Println(fruit)
	}
}

func printColor(c <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for color := range c {
		fmt.Println(color)
	}
}

func main() {
	// Creating a slice with length 3 and capacity 5
	numbers := make([]int, 3, 5)

	// Adding elements to the slice
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	// Appending elements beyond the length
	numbers = append(numbers, 4, 5, 6, 7, 8)

	// Printing length and capacity
	fmt.Printf("Slice: %v\n", numbers)
	fmt.Printf("Length: %d\n", len(numbers))
	fmt.Printf("Capacity: %d\n", cap(numbers))
}
