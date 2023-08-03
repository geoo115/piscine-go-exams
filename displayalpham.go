package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for ch := 'a'; ch <= 'z'; ch++ {
		if ch%2 == 0 {
			z01.PrintRune(ch - 32) // Uppercase even letters
		} else {
			z01.PrintRune(ch) // Lowercase odd letters
		}
	}
	z01.PrintRune('\n')
}
