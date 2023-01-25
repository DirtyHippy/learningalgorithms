package main

func computeMinRefills(dist, tank int, stops []int) int {
	var refills int

	stops = append([]int{0}, stops...)
	stops = append(stops, dist)
	var fuel = tank

	for i := 1; i < len(stops); {
		if fuel-(stops[i]-stops[i-1]) >= 0 {
			fuel -= stops[i] - stops[i-1]
			i++
		} else if fuel == tank {
			refills = -1
			break
		} else {
			fuel = tank
			refills++
		}
	}

	return refills
}
