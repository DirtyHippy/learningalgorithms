package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	table := make(map[int]int, amount+1)
	table[0] = 0
	for i := 1; i < amount+1; i++ {
		table[i] = math.MaxInt
	}

	for m := 1; m <= amount; m++ {
		found := false
		for _, c := range coins {
			if c <= m {
				if table[m-c] != -1 {
					table[m] = int(math.Min(float64(table[m]), float64(1+table[m-c])))
					found = true
				}
			}
		}
		if !found {
			table[m] = -1
		}
	}

	return table[amount]
}

func main() {
	fmt.Println(coinChange([]int{2}, 1))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{2}, 4))
	fmt.Println(coinChange([]int{2, 4}, 5))
}
