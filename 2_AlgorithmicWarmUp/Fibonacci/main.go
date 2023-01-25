package main

import "fmt"

var cache = make(map[int]int)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Fibonacci(n))
}

func Fibonacci(n int) int {
	if _, ok := cache[n]; !ok {
		if n <= 1 {
			cache[n] = n
		} else {
			cache[n] = Fibonacci(n-2) + Fibonacci(n-1)
		}
	}

	return cache[n]
}
