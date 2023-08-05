package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isNotDuplicate(char rune, str string) bool {
	for _, c := range str {
		if char == c {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	a := os.Args[1]
	b := os.Args[2]
	var result string

	for _, char := range a {
		if isNotDuplicate(char, result) {
			result += string(char)
		}
	}

	for _, char := range b {
		if isNotDuplicate(char, result) {
			result += string(char)
		}
	}

	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
