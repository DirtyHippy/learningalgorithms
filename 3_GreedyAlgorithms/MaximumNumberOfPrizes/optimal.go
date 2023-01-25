package main

import (
	"fmt"
)

func optimalSummands(n int) []int {
	var summands []int = make([]int, 0)

	for i := 1; i <= n; i++ {
		n -= i
		summands = append(summands, i)
	}

	summands[len(summands)-1] += n

	return summands
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	summands := optimalSummands(n)
	fmt.Println(len(summands))
	for _, summand := range summands {
		fmt.Printf("%d ", summand)
	}
}
