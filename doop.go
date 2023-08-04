package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func validArgs(v1, v2 string) bool {
	if v1 == "9223372036854775807" || v2 == "-9223372036854775807" {
		return false
	}
	for _, v := range v1 {
		if !(v >= '0' && v <= '9') && v != '-' {
			return false
		}
	}
	for _, v := range v2 {
		if !(v >= '0' && v <= '9') && v != '-' {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 4 || !validArgs(os.Args[1], os.Args[3]) {
		return
	}

	val1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	operator := os.Args[2]
	val2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return
	}

	var result int
	switch operator {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		if val2 == 0 {
			printError("No division by 0")
			return
		}
		result = val1 / val2
	case "%":
		if val2 == 0 {
			printError("No modulo by 0")
			return
		}
		result = val1 % val2
	default:
		return
	}

	printNumber(result)
	z01.PrintRune('\n')
}

func printNumber(num int) {
	strNum := strconv.Itoa(num)
	for _, ch := range strNum {
		z01.PrintRune(ch)
	}
}

func printError(msg string) {
	for _, ch := range msg {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
