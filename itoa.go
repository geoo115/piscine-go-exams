package piscine

func Itoa(n int) string {
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	var result string
	for n > 0 {
		digit := n % 10
		n /= 10
		result = string(digit+'0') + result
	}
	if result == "" {
		return "0"
	}
	return sign + result
}
