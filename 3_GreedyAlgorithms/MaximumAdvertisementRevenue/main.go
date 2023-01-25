package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	clicks := make([]int, n)
	prices := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&clicks[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}

	fmt.Print(MaximumAdvertisementRevenue(clicks, prices))
}

func MaximumAdvertisementRevenue(clicks, prices []int) int {
	maxSum := 0
	sort.Sort(sort.IntSlice(clicks))
	sort.Sort(sort.IntSlice(prices))
	for i := 0; i < len(clicks); i++ {
		maxSum += clicks[i] * prices[i]
	}
	return maxSum
}
