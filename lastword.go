package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	lastspace := len(input) - 1

	// Move backward to find the start of the last word
	for lastspace >= 0 && input[lastspace] == ' ' {
		lastspace--
	}
	if lastSpace < 0 {
		return
	}
	end := lastspace

	// Move backward to find the end of the last word
	for lastspace >= 0 && input[lastspace] != ' ' {
		lastspace--
	}

	// Print the last word
	if lastspace >= 0 {
		for i := lastspace + 1; i <= end; i++ {
			z01.PrintRune(rune(input[i]))
		}
		z01.PrintRune('\n')
	}
}
