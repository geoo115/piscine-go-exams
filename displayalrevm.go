package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for ch := 'Z'; ch >= 'A'; ch-- {
		if ch%2 == 0 {
			z01.PrintRune(ch + 32)
		} else {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
