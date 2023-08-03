package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	var word string
	var expand string
	for _, ch := range os.Args[1] {
		if ch == ' ' {
			if len(word) > 0 {
				expand += word + "   "
				word = ""
			}
		} else {
			word += string(ch)
		}
	}
	if len(word) > 0 {
		expand += word
	}
	for _, r := range expand {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
