package main

import (
	"fmt"
	"unicode"
)

func CountWords(s string) int {
	fmt.Println(s)
	count := 1
	for _, letter := range s {
		if unicode.IsUpper(letter) {
			count += 1
		}
	}
	return count
}

func main() {
	count := CountWords("saveChangesInTheEditor")
	fmt.Println("Word count : ", count)
}
