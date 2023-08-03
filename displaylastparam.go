package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args

	for _, ch := range input[len(input)-1] {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
