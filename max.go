package piscine

func Max(a []int) int {
	b := 0
	for _, c := range a {
		if c > b {
			b = c
		}
	}
	return b
}
