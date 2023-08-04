package piscine

import (
	"github.com/01-edu/z01"
)

func isAlphaNumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func Capitalize(s string) string {
	capitalizeNext := true
	var result string

	for _, ch := range s {
		if isAlphaNumeric(ch) {
			if capitalizeNext {
				if ch >= 'a' && ch <= 'z' {
					ch -= 'a' - 'A' // Convert to uppercase
				}
				capitalizeNext = false
			} else {
				if ch >= 'A' && ch <= 'Z' {
					ch += 'a' - 'A' // Convert to lowercase
				}
			}
			z01.PrintRune(ch)
			result += string(ch)
		} else {
			capitalizeNext = true
			z01.PrintRune(ch)
			result += string(ch)
		}
	}

	return result
}
