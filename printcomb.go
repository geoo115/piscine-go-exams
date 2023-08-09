package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := '1'; j <= '8'; j++ {
			for x := '2'; x <= '9'; x++ {
				if i != j && i != x && j != x && x > j && j > i {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(x)
					if j != x && i == '7' && j == '8' && x == '9' {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
