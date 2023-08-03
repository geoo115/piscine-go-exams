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

	if len(a) > len(b) {
		return
	}

	// Two pointers to keep track of the positions in both strings
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}

	if i == len(a) {
		for _, ch := range a {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
