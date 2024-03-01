package main

import (
	"fmt"
)

func main() {
	//sentence := "india is a beautiful coutry, india  usa uk uk india uk usa spain usa spain usa"

	s := "aaabbddeerrraa"
	count := 1
	current := string(s[0])
	length := len(s)
	index := 1
	for index < length {
		if string(s[index]) == current {
			count++
		} else {
			fmt.Print(string(s[index-1]), count)
			count = 1
			current = string(s[index])
		}
		index++
	}
	if count > 0 {
		fmt.Print(string(s[index-1]), count)
	}
	fmt.Println()
}
