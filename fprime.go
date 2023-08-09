package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func printPrimeFactors(num int) {
	originalNum := num
	isFirstFactor := true

	for i := 2; i <= num; {
		if num%i == 0 && isPrime(i) {
			if !isFirstFactor {
				z01.PrintRune('*')
			}
			for _, r := range strconv.Itoa(i) {
				z01.PrintRune(r)
			}
			num /= i
			isFirstFactor = false
		} else {
			i++
		}
	}

	if originalNum != num {
		z01.PrintRune('\n')
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil || num <= 1 {
		return
	}

	printPrimeFactors(num)
}
