package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	var rot13 string
	for _, ch := range os.Args[1] {
		if ch >= 'a' && ch <= 'z' {
			rot13 += string((ch+13-'a')%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			rot13 += string((ch+13-'A')%26 + 'A')
		} else {
			rot13 += string(ch)
		}
	}
	for _, ch := range rot13 {
		z01.PrintRune(ch)
	}

	z01.PrintRune('\n')
}
