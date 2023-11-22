package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(char rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if char == v {
			return true
		}
	}
	return false
}

func pigLatin(s string) string {

	if isVowel(rune(s[0])) {
		return s + "ay"
	}

	firstVowelIndex := -1
	for i, char := range s {
		if isVowel(char) {
			firstVowelIndex = i
			break
		}
	}

	if firstVowelIndex == -1 {
		return "No vowels"
	}

	return s[firstVowelIndex:] + s[:firstVowelIndex] + "ay"
}

func main() {
	if len(os.Args) != 2 {
		return
	}
		if len(os.Args[1]) == 0 {
		return
	}

	result := pigLatin(os.Args[1])
	for _, c := range result {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
