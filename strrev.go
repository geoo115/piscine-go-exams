package piscine

func StrRev(s string) string {
	var rev string
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}
