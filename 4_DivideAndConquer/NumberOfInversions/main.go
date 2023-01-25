package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}

	fmt.Println(NumberOfInversions(input))
}

func NumberOfInversions(input []int) int {
	numInversion := 0
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				numInversion++
			}
		}
	}
	return numInversion
}
