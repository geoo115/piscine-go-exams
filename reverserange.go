package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	start, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	end, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	if start <= end {
		for i := end; i >= start; i-- {
			fmt.Printf("%d ", i)
		}
	} else {
		for i := end; i <= start; i++ {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
