package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				if i != j && j != k && i != k {
					z01.PrintRune(rune('0' + i))
					z01.PrintRune(rune('0' + j))
					z01.PrintRune(rune('0' + k))
					if i != 2 || j != 1 || k != 0 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
