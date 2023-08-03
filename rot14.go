package piscine

func Rot14(s string) string {
	var rot14 string
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			rot14 += string((ch+14-'a')%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			rot14 += string((ch+14-'A')%26 + 'A')
		} else {
			rot14 += string(ch)
		}
	}
	return rot14
}
