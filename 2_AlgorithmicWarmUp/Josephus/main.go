package main

import "fmt"

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	fmt.Println(Josephus(n, k))
}

func Josephus(n, k int) int {
	lives := make([]bool, 0, n)
	for i := 0; i < n; i++ {
		lives = append(lives, true)
	}

	numberOfSurvivors := n
	currentPosition := 0
	index := 0

	for {
		if lives[currentPosition] {
			index++
		}
		if index == k {
			numberOfSurvivors--
			if numberOfSurvivors == 0 {
				return currentPosition
			} else {
				lives[currentPosition] = false
				index = 0
			}
		}
		currentPosition = (currentPosition + 1) % n
	}
}
