package main

import "fmt"

func SliceDifference(first, second []string) []string {
	excluded := make(map[string]struct{}, len(second))
	for _, item := range second {
		excluded[item] = struct{}{}
	}

	result := make([]string, 0, len(first))
	for _, item := range first {
		if _, found := excluded[item]; !found {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(SliceDifference(slice1, slice2))
}
