package main

import "fmt"

func pivotIndex(nums []int) int {
	sumLeft := make([]int, len(nums))
	sumRight := make([]int, len(nums))
	index := -1

	sumLeft[0] = 0
	sumRight[len(nums)-1] = 0

	for i := 1; i < len(nums); i++ {
		sumLeft[i] = sumLeft[i-1] + nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		sumRight[i] = sumRight[i+1] + nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		if sumLeft[i] == sumRight[i] {
			index = i
			break
		}
	}

	return index
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}
