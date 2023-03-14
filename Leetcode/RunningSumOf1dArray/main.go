package main

import "fmt"

func runningSum(nums []int) []int {
	out := make([]int, len(nums))
	out[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		out[i] = out[i-1] + nums[i]
	}

	return out
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}
