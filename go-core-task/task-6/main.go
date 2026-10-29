package main

import (
	"fmt"
	"math/rand"
)

func RandomGenerator() <-chan int {
	ch := make(chan int)

	go func() {
		for {
			ch <- rand.Intn(100)
		}
	}()

	return ch
}

func Take(ch <-chan int, n int) []int {
	result := make([]int, 0, n)
	for range n {
		result = append(result, <-ch)
	}
	return result
}

func main() {
	ch := RandomGenerator()
	for _, v := range Take(ch, 10) {
		fmt.Println(v)
	}
}
