package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var n int

	fmt.Scan(&n)

	a := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(MaxPairwiseProductFastOneLoop(a))
}

func MaxPairwiseProductFastOneLoop(a []int) int {
	maxVal := 0
	secondMaxVal := 0
	for _, val := range a {
		if val > maxVal {
			secondMaxVal, maxVal = maxVal, val
		} else if val > secondMaxVal {
			secondMaxVal = val
		}
	}
	return secondMaxVal * maxVal
}

func MaxPairwiseProductFast(a []int) int {
	maxValIndex := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxValIndex] {
			maxValIndex = i
		}
	}

	a[len(a)-1], a[maxValIndex] = a[maxValIndex], a[len(a)-1]
	maxValIndex = len(a) - 1

	secondMaxValIndex := 0
	for i := 1; i < len(a)-1; i++ {
		if a[i] > a[secondMaxValIndex] {
			secondMaxValIndex = i
		}
	}
	return a[secondMaxValIndex] * a[maxValIndex]
}

func MaxPairwiseProductNaive(a []int) int {
	var maxVal int

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if i != j {
				res := a[i] * a[j]
				if res > maxVal {
					maxVal = res
				}
			}
		}
	}

	return maxVal
}

func MaxPairwiseProductBySorting(a []int) int {
	sort.Sort(sort.IntSlice(a))

	return a[len(a)-1] * a[len(a)-2]
}

func StressTest(n, m int) {
	for {
		n = rand.Intn(n-2) + 2
		a := make([]int, 0, n)
		for i := 0; i < len(a); i++ {
			a[i] = rand.Intn(m)
		}

		naive := MaxPairwiseProductNaive(a)
		fast := MaxPairwiseProductFast(a)

		if naive != fast {
			fmt.Println("Wrong answer:", naive, fast)
		}
	}
}
