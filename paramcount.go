package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(rune(len(os.Args) - 1 + '0'))
	z01.PrintRune('\n')
}
