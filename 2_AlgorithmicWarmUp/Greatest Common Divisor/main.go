package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(GreatestCommonDivisor(a, b))
}

func GreatestCommonDivisor(a, b int) int {
	for a > 0 && b > 0 {
		if a >= b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a >= b {
		return a
	} else {
		return b
	}
}
