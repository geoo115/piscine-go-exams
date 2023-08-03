package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	var result string
	for _, ch := range os.Args[1] {
		if ch >= 'a' && ch <= 'z' {
			result += string(ch - 32)
		} else if ch >= 'A' && ch <= 'Z' {
			result += string(ch + 32)
		} else {
			result += string(ch)
		}
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
