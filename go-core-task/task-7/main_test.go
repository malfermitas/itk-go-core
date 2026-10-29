package main

import (
	"context"
	"sort"
	"testing"
	"time"
)

const testTimeout = 2 * time.Second

func intChannels(values ...[]int) []chan any {
	channels := make([]chan any, 0, len(values))
	for _, vals := range values {
		ch := make(chan any, len(vals))
		for _, v := range vals {
			ch <- v
		}
		close(ch)
		channels = append(channels, ch)
	}
	return channels
}

func collectInts(t *testing.T, ch <-chan any) []int {
	t.Helper()
	timeout := time.After(testTimeout)
	var got []int
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return got
			}
			n, ok := v.(int)
			if !ok {
				t.Fatalf("unexpected value type: %T", v)
			}
			got = append(got, n)
		case <-timeout:
			t.Fatal("timed out waiting for merged channel to close")
			return got
		}
	}
}

func TestMergeSingleChannel(t *testing.T) {
	channels := intChannels([]int{1, 2, 3})

	got := collectInts(t, merge[any](context.Background(), channels...))

	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestMergeMultipleChannels(t *testing.T) {
	channels := intChannels(
		[]int{1, 2, 3},
		[]int{4, 5},
		[]int{6},
	)

	got := collectInts(t, merge[any](context.Background(), channels...))

	want := []int{1, 2, 3, 4, 5, 6}
	sort.Ints(got)
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestMergeNoChannels(t *testing.T) {
	got := collectInts(t, merge[any](context.Background()))

	if len(got) != 0 {
		t.Fatalf("got %v, want empty result", got)
	}
}

func TestMergeClosesWhenAllChannelsClosed(t *testing.T) {
	channels := intChannels([]int{42})

	merged := merge[any](context.Background(), channels...)
	collectInts(t, merged)

	select {
	case _, ok := <-merged:
		if ok {
			t.Fatal("expected channel to be closed")
		}
	case <-time.After(testTimeout):
		t.Fatal("timed out waiting for channel to close")
	}
}

func TestMergeManyChannelsNoDataLoss(t *testing.T) {
	const (
		channelCount = 10
		valueCount   = 100
	)

	channels := make([]chan any, channelCount)
	for i := 0; i < channelCount; i++ {
		ch := make(chan any, valueCount)
		for j := 0; j < valueCount; j++ {
			ch <- i*valueCount + j
		}
		close(ch)
		channels[i] = ch
	}

	got := collectInts(t, merge[any](context.Background(), channels...))

	total := channelCount * valueCount
	if len(got) != total {
		t.Fatalf("got %d values, want %d", len(got), total)
	}

	sort.Ints(got)
	for i, v := range got {
		if v != i {
			t.Fatalf("unexpected value at index %d: got %d, want %d", i, v, i)
		}
	}
}

func TestMergeStopsOnContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan any)

	merged := merge[any](ctx, ch)
	cancel()

	select {
	case _, ok := <-merged:
		if ok {
			t.Fatal("expected channel to be closed after cancellation")
		}
	case <-time.After(testTimeout):
		t.Fatal("timed out waiting for channel to close after cancellation")
	}
}
