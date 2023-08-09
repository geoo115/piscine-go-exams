package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(a) == 0 {
		return nil
	}

	var start, end int

	switch len(nbrs) {
	case 0:
		return a
	case 1:
		start = nbrs[0]
		end = len(a)
	default:
		start = nbrs[0]
		end = nbrs[1]
	}

	if start < 0 {
		start = len(a) + start
	}
	if end < 0 {
		end = len(a) + end
	}

	if start < 0 {
		start = 0
	} else if start > len(a) {
		start = len(a)
	}

	if end < 0 {
		end = 0
	} else if end > len(a) {
		end = len(a)
	}

	if start > end {
		return nil
	}

	if start == end {
		return []string{}
	}

	return a[start:end]
}
