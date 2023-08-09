package piscine

func ForEach(f func(int), a []int) {
	for _, num := range a {
		f(num)
	}
}
