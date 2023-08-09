package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	var words []string
	var word string
	for _, r := range str {
		if r == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
			continue
		}
		word += string(r)
	}
	if word != "" {
		words = append(words, word)
	}
	for i := len(words) - 1; i >= 0; i-- {
		if i != len(words)-1 {
			z01.PrintRune(' ')
		}
		for _, r := range words[i] {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
