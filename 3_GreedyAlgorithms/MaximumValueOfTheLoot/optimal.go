package main

import (
	"fmt"
	"sort"
)

type Compound struct {
	value   int
	weights int
}

func getOptimalValue(capacity int, compounds []Compound) float64 {
	var value float64

	sort.Slice(compounds, func(i, j int) bool {
		return float64(compounds[i].value)/float64(compounds[i].weights) >
			float64(compounds[j].value)/float64(compounds[j].weights)
	})

	for capacity > 0 && len(compounds) > 0 {
		availableWeight := capacity
		if capacity > compounds[0].weights {
			availableWeight = compounds[0].weights
		}

		var cost float64 = float64(compounds[0].value) / float64(compounds[0].weights)
		value += cost * float64(availableWeight)
		capacity -= availableWeight
		compounds = compounds[1:]
	}

	return value
}

func main() {
	var n, capacity int
	fmt.Scanf("%d %d", &n, &capacity)

	compounds := make([]Compound, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &compounds[i].value, &compounds[i].weights)
	}

	fmt.Printf("%.4f\n", getOptimalValue(capacity, compounds))
}
