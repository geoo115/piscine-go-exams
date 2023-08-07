package piscine

import (
	"strconv"

	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	var result int
	for i := 0; i < len(a)-1; i++ {
		if i == 0 {
			result = a[0]
		}
		result = f(result, a[i+1])
	}
	for _, char := range strconv.Itoa(result) {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n)
}
