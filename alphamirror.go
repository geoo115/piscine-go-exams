package main

import (
	"os"

	"github.com/01-edu/z01"
)

func alphamirror(input string) string {
	var result string

	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			opposite := 'Z' - (ch - 'A')
			result += string(opposite)
		} else if ch >= 'a' && ch <= 'z' {
			opposite := 'z' - (ch - 'a')
			result += string(opposite)
		} else {
			result += string(ch)
		}
	}

	return result
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	result := alphamirror(input)

	for _, ch := range result {
		z01.PrintRune(ch)
	}

	z01.PrintRune('\n')
}
