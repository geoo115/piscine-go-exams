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
	lastSpace := len(input) - 1

	for lastSpace >= 0 && input[lastSpace] != ' ' {
		lastSpace--
	}
	for i := lastSpace + 1; i < len(input); i++ {
		z01.PrintRune(rune(input[i]))
	}
	z01.PrintRune('\n')
}
