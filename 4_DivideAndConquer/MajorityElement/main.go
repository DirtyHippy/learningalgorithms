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

	fmt.Println(MajorityElement(input))
}

func MajorityElement(input []int) int {
	counter := make(map[int]int)
	for _, val := range input {
		counter[val] += 1
	}

	for _, val := range counter {
		if val > len(input)/2 {
			return 1
		}
	}
	return 0
}
