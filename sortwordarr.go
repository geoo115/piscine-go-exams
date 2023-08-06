package piscine

import (
	"github.com/01-edu/z01"
)

func SortWordArr(array []string) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
