package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	a := os.Args[1]
	b := os.Args[2]
	var result string

	for _, char := range a {
		for _, bChar := range b {
			if char == bChar {
				result += string(char)
				break
			}
		}
	}

	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
