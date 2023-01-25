package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Println(maximumSalary(numbers))
}

func maximumSalary(numbers []int) string {
	maxSalary := ""
	for len(numbers) != 0 {
		bestNumber := 0
		bestNumberIndex := 0
		for i := 0; i < len(numbers); i++ {
			if isBetter(numbers[i], bestNumber) {
				bestNumber = numbers[i]
				bestNumberIndex = i
			}
		}

		maxSalary += strconv.Itoa(bestNumber)

		numbers[bestNumberIndex] = numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
	}

	return maxSalary
}

func isBetter(n1, n2 int) bool {
	salary1, _ := strconv.Atoi(strconv.Itoa(n1) + strconv.Itoa(n2))
	salary2, _ := strconv.Atoi(strconv.Itoa(n2) + strconv.Itoa(n1))

	return salary1 > salary2
}
