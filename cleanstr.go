package main

import (
	"os"

	"github.com/01-edu/z01"
)

func cleanString(s string) string {
	var words []string
	wordStart := false
	word := ""
	for _, r := range s {
		if r == ' ' || r == '\t' {
			if wordStart {
				words = append(words, word)
				word = ""
			}
			wordStart = false
		} else {
			word += string(r)
			wordStart = true
		}
	}
	if wordStart {
		words = append(words, word)
	}
	return joinWords(words)
}

func joinWords(words []string) string {
	var result string
	for i, word := range words {
		result += word
		if i < len(words)-1 {
			result += " "
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	result := cleanString(os.Args[1])
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
