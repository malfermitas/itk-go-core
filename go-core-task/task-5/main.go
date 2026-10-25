package main

import "fmt"

func findIntersections(a, b []int) (bool, []int) {
	setA := make(map[int]struct{}, len(a))
	for _, v := range a {
		setA[v] = struct{}{}
	}

	intersections := make([]int, 0)
	seen := make(map[int]struct{})

	for _, v := range b {
		if _, ok := setA[v]; !ok {
			continue
		}
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		intersections = append(intersections, v)
	}

	return len(intersections) > 0, intersections
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, intersections := findIntersections(a, b)
	fmt.Println(hasIntersection, intersections)
}
