package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// gcd calculates the greatest common divisor of two integers using the Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Check if there are exactly 2 command-line arguments
	if len(os.Args) != 3 {
		return
	}

	// Parse the two input strings to integers
	a, errA := strconv.Atoi(os.Args[1])
	b, errB := strconv.Atoi(os.Args[2])

	// Check if the parsing was successful
	if errA != nil || errB != nil {
		return
	}

	// Calculate the GCD of the two integers
	result := gcd(a, b)

	// Print the result using z01.PrintRune
	for _, char := range strconv.Itoa(result) {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
