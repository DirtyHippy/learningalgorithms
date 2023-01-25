package main

import (
	"fmt"
)

func main() {
	var totalDistance, capacity, n int
	fmt.Scan(&totalDistance, &capacity, &n)

	stops := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&stops[i])
	}

	fmt.Print(Refills(totalDistance, capacity, stops))
}

func Refills(totalDistance, capacity int, stops []int) int {
	stops = append([]int{0}, stops...)
	stops = append(stops, totalDistance)
	numRefills := 0
	remainingCapacity := capacity

	if capacity >= stops[len(stops)-1] {
		return 0
	}

	for i := 1; i < len(stops); {
		curStop, prevStop := stops[i], stops[i-1]
		if curStop <= prevStop || curStop-prevStop > capacity {
			return -1
		} else {
			diff := curStop - prevStop
			if remainingCapacity-diff >= 0 {
				remainingCapacity -= diff
				i++
			} else {
				remainingCapacity = capacity
				numRefills++
			}
		}
	}

	return numRefills
}
