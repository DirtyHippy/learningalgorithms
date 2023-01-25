package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(FibonacciLastDigit(n))
}

func FibonacciLastDigit(n int) int {
	if n <= 1 {
		return n
	}

	f := make([]int, n+1, n+1)
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2]) % 10
	}
	return f[n]
}
