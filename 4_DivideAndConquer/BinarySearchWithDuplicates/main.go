package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
	fmt.Scan(&n)
	searchData := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&searchData[i])
	}
	fmt.Println(strings.Trim(fmt.Sprint(BinarySearch(input, searchData)), "[]"))
}

func BinarySearch(input, searchData []int) []int {
	output := make([]int, 0, len(searchData))
	for _, val := range searchData {
		minIndex := 0
		maxIndex := len(input) - 1
		result := -1
		for maxIndex >= minIndex {
			middleIndex := (minIndex + maxIndex) / 2
			if input[middleIndex] == val {
				maxIndex = middleIndex - 1
				result = middleIndex
			} else if input[middleIndex] < val {
				minIndex = middleIndex + 1
			} else {
				maxIndex = middleIndex - 1
			}
		}
		output = append(output, result)
	}
	return output
}
