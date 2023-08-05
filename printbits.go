package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func printBits(num int) {
	for i := 7; i >= 0; i-- {
		bit := (num >> uint(i)) & 1
		if bit == 1 {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
	}
}

func main() {
	if len(os.Args) != 2 {

		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	printBits(num)
	z01.PrintRune('\n')
}
