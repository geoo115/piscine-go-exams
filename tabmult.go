package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	for i := 1; i <= 9; i++ {
		result := i * num
		PrintNum(i)
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		PrintNum(num)
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		PrintNum(result)
		z01.PrintRune('\n')
	}
}

func PrintNum(num int) {
	// Convert the number to a string
	strNum := strconv.Itoa(num)

	// Print the characters of the string one by one
	for _, ch := range strNum {
		z01.PrintRune(ch)
	}
}
