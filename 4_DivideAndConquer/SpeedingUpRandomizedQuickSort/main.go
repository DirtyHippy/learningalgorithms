package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}

	fmt.Println(strings.Trim(fmt.Sprint(QuickSort(input)), "[]"))
}

func QuickSort(input []int) []int {
	if len(input) == 0 || len(input) == 1 {
		return input
	}
	random := input[rand.Intn(len(input))]
	small, large, equal := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, val := range input {
		if val < random {
			small = append(small, val)
		} else if val > random {
			large = append(large, val)
		} else {
			equal = append(equal, val)
		}
	}
	small = QuickSort(small)
	large = QuickSort(large)
	return append(small, append(equal, large...)...)
}
