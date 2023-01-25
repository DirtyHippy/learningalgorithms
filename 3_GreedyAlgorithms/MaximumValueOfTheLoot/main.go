package main

import (
	"fmt"
	"math"
)

func main() {
	var n, capacity int
	fmt.Scan(&n, &capacity)
	costs := make([]int, n, n)
	weights := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&costs[i], &weights[i])
	}
	fmt.Println(MaximumValueOfTheLoot(costs, weights, float64(capacity)))
}

func MaximumValueOfTheLoot(costs, weights []int, capacity float64) float64 {
	if capacity == 0 || len(weights) == 0 {
		return 0
	}
	maxPriceIndex, maxPrice := 0, 0.0

	for i, _ := range weights {
		price := float64(costs[i]) / float64(weights[i])
		if price > maxPrice {
			maxPriceIndex = i
			maxPrice = price
		}
	}
	amount := math.Min(float64(weights[maxPriceIndex]), capacity)
	value := float64(costs[maxPriceIndex]) * amount / float64(weights[maxPriceIndex])

	costs[maxPriceIndex] = costs[len(costs)-1]
	costs = costs[:len(costs)-1]

	weights[maxPriceIndex] = weights[len(weights)-1]
	weights = weights[:len(weights)-1]

	return value + MaximumValueOfTheLoot(costs, weights, capacity-amount)
}
