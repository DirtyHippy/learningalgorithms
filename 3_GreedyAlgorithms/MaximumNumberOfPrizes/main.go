package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	prizes := MaximumNumberOfPrizes(n)
	fmt.Println(len(prizes))
	fmt.Println(strings.Trim(fmt.Sprint(prizes), "[]"))
}

func MaximumNumberOfPrizes(n int) []int {
	k := 1
	for i := k; i <= n; i++ {
		if n < i*(i+1)/2 {
			k = i
			break
		}
	}
	k--
	res := make([]int, 0, k)
	delta := n - k*(k+1)/2
	for i := 1; i < k; i++ {
		res = append(res, i)
	}
	res = append(res, k+delta)
	return res
}
