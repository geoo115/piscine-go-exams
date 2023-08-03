package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, ch := range os.Args[1] {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
