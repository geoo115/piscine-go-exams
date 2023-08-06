package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func addPrimesSum(limit int) int {
	sum := 0
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func printNumber(num int) {
	if num >= 10 {
		printNumber(num / 10)
	}
	z01.PrintRune(rune(num%10 + '0'))
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	result := addPrimesSum(num)
	printNumber(result)
	z01.PrintRune('\n')
}
