package piscine

func Isalpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func ReverseStrCap(s string) string {
	cap := true
	var result string

	runes := []rune(s)

	for i := len(runes) - 1; i >= 0; i-- {
		ch := runes[i]

		if Isalpha(ch) {
			if cap {
				if ch >= 'a' && ch <= 'z' {
					ch = ch - 32
				}
				cap = false
			} else {
				if ch >= 'A' && ch <= 'Z' {
					ch = ch + 32
				}
			}
			result = string(ch) + result
		} else {
			cap = true
			result = string(ch) + result
		}
	}

	return result
}
