package main

import (
	"fmt"
	"math"
)

func main() {
	var money int
	fmt.Scan(&money)
	fmt.Println(change(money))
}

func changeRecursive(money int) int {
	if money == 0 {
		return 0
	}
	result := math.MaxInt

	for _, c := range [3]int{1, 3, 4} {
		if c <= money {
			result = int(math.Min(float64(result), float64(1+changeRecursive(money-c))))
		}
	}

	return result
}

func change(money int) int {
	table := make(map[int]int, money+1)
	table[0] = 0
	for i := 1; i < money+1; i++ {
		table[i] = math.MaxInt
	}

	for m := 1; m <= money; m++ {
		for _, c := range [3]int{1, 3, 4} {
			if c <= m {
				table[m] = int(math.Min(float64(table[m]), float64(1+table[m-c])))
			}
		}
	}
	return table[money]
}
