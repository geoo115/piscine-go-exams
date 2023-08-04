package piscine

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune(' ')
		z01.PrintRune('\n')
		return
	}
	if size > len(slice) {
		z01.PrintRune('[')
		z01.PrintRune(']')
		z01.PrintRune('\n')
		return
	}

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}

		// Print the chunk as a slice
		z01.PrintRune('[')
		for j := i; j < end; j++ {
			z01.PrintRune(rune('0' + slice[j]))
			if j < end-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')

		if i+size < len(slice) {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
