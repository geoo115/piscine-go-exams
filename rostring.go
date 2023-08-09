package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	input := args[0]

	startIdx := 0
	for startIdx < len(input) && isSpace(rune(input[startIdx])) {
		startIdx++
	}

	endIdx := len(input) - 1
	for endIdx >= 0 && isSpace(rune(input[endIdx])) {
		endIdx--
	}

	if startIdx > endIdx {
		z01.PrintRune('\n')
		return
	}

	// Find the index of the first space after the first word
	spaceIdx := startIdx
	for spaceIdx <= endIdx && !isSpace(rune(input[spaceIdx])) {
		spaceIdx++
	}

	// Print words after the first word
	for i := spaceIdx; i <= endIdx; i++ {
		if isSpace(rune(input[i])) {
			continue
		}
		wordStart := i
		for wordStart <= endIdx && !isSpace(rune(input[wordStart])) {
			wordStart++
		}
		for j := i; j < wordStart; j++ {
			z01.PrintRune(rune(input[j]))
		}
		z01.PrintRune(' ')
		i = wordStart
	}

	// Print the first word
	for i := startIdx; i < spaceIdx; i++ {
		z01.PrintRune(rune(input[i]))
	}

	z01.PrintRune('\n')
}
