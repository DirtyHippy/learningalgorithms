package main

import "testing"

type testValues struct {
	totalDistance int
	capacity      int
	stops         []int
	numRefills    int
}

var tests = []testValues{
	{100, 3, []int{1, 2, 5, 9}, -1},
	{200, 250, []int{100, 150}, 0},
	{950, 400, []int{200, 375, 550, 750}, 2},
}

func TestAverage(t *testing.T) {
	for _, val := range tests {
		res := Refills(val.totalDistance, val.capacity, val.stops)
		if res != val.numRefills {
			t.Error(
				"For", val.stops,
				"expected", val.numRefills,
				"got", res,
			)
		}
	}
}
