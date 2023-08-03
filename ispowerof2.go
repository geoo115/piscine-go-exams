package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	if num < 0 {
		return
	}
	if IspowerOf2(num) {
		for _, r := range "true\n" {
			z01.PrintRune(r)
		}

	} else {
		for _, r := range "false\n" {
			z01.PrintRune(r)
		}
	}

}

func IspowerOf2(n int) bool {
	if n == 0 {
		return true
	}
	if n%2 == 0 {
		n /= 2
	}
	return n == 1
}
