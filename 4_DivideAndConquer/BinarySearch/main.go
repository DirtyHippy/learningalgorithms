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
	output := make([]int, 0, len(input))
	for _, val := range searchData {
		minIndex := 0
		maxIndex := len(input) - 1
		found := false
		for maxIndex >= minIndex {
			middleIndex := (minIndex + maxIndex) / 2
			if input[middleIndex] == val {
				found = true
				output = append(output, middleIndex)
				break
			} else if input[middleIndex] < val {
				minIndex = middleIndex + 1
			} else {
				maxIndex = middleIndex - 1
			}
		}
		if !found {
			output = append(output, -1)
		}
	}
	return output
}
