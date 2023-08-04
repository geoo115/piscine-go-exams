package piscine

import "github.com/01-edu/z01"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, val := range a {
		n = f(n, val)
	}
	printNumber(n)
	z01.PrintRune('\n')
}

func printNumber(num int) {
	if num == 0 {
		z01.PrintRune('0')
		return
	}

	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}

	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune('0' + digits[i]))
	}
}
