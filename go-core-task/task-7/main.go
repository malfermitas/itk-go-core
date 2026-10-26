package main

import (
	"context"
	"fmt"
	"sync"
)

func merge[T any](ctx context.Context, chans ...chan T) chan T {
	resultChan := make(chan T)
	wg := sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)
		go func(ch chan T) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-ch:
					if !ok {
						return
					}
					select {
					case resultChan <- val:
					case <-ctx.Done():
						return
					}
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

func main() {
	ctx := context.Background()

	const n = 5
	channels := make([]chan any, 0, n)
	for i := 0; i < n; i++ {
		ch := make(chan any, i+1)
		for j := 0; j <= i; j++ {
			ch <- fmt.Sprintf("channel %d, value %d", i, j)
		}
		close(ch)
		channels = append(channels, ch)
	}

	merged := merge[any](ctx, channels...)
	for value := range merged {
		fmt.Println(value)
	}
}
