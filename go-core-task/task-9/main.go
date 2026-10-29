package main

import "fmt"

func cube(n uint8) float64 {
	f := float64(n)
	return f * f * f
}

func process(in <-chan uint8, out chan<- float64) {
	defer close(out)
	for n := range in {
		out <- cube(n)
	}
}

func main() {
	numbers := make(chan uint8)
	results := make(chan float64)

	go process(numbers, results)

	go func() {
		for i := uint8(1); i <= 10; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
