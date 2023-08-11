package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	a := os.Args[1]
	b := os.Args[2]
	c := os.Args[3]
	var result string
	for _, ch := range a {
		if ch == rune(b[0]) {
			ch = rune(c[0])
		}
		result += string(ch)
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
