package main

import "testing"

type testValues struct {
	clicks []int
	prices []int
	maxSum int
}

var tests = []testValues{
	{clicks: []int{23}, prices: []int{39}, maxSum: 897},
	{clicks: []int{2, 3, 9}, prices: []int{7, 4, 2}, maxSum: 79},
}

func TestAverage(t *testing.T) {
	for _, val := range tests {
		res := MaximumAdvertisementRevenue(val.clicks, val.prices)
		if res != val.maxSum {
			t.Error(
				"expected", val.maxSum,
				"got", res,
			)
		}
	}
}
