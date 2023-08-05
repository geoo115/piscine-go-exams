package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	for _, ch := range os.Args[1] {
		z01.PrintRune(ch)
		rep := int(ch - 'a')
		for i := 0; i < rep; i++ {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
