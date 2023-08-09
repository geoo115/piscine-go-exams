package piscine

func AtoiBase(s string, base string) int {
	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		baseMap[base[i]] = i
	}

	result := 0
	sign := 1
	baseLength := len(base)
	sLength := len(s)

	for i := 0; i < sLength; i++ {
		char := s[i]
		if i == 0 && (char == '-' || char == '+') {
			if char == '-' {
				sign = -1
			}
			continue
		}
		value, exists := baseMap[char]
		if !exists {
			return 0
		}
		result = result*baseLength + value
	}

	return result * sign
}
