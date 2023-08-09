package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func printError(message string) {
	for _, r := range message {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	start, err := strconv.Atoi(os.Args[1])
	if err != nil {
		printError(err.Error() + "\n")
		return
	}

	end, err := strconv.Atoi(os.Args[2])
	if err != nil {
		printError(err.Error() + "\n")
		return
	}

	if start <= end {
		for i := start; i <= end; i++ {
			for _, r := range strconv.Itoa(i) {
				z01.PrintRune(r)
			}
			z01.PrintRune(' ')
		}
	} else {
		for i := start; i >= end; i-- {
			for _, r := range strconv.Itoa(i) {
				z01.PrintRune(r)
			}
			z01.PrintRune(' ')
		}
	}
}
