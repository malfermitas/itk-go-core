package main

import (
	"testing"
	"time"
)

func TestCube(t *testing.T) {
	tests := []struct {
		name string
		in   uint8
		want float64
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 8},
		{"three", 3, 27},
		{"ten", 10, 1000},
		{"max", 255, 16581375},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cube(tt.in); got != tt.want {
				t.Fatalf("cube(%d) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestProcess(t *testing.T) {
	input := []uint8{0, 1, 2, 3, 10, 255}
	want := []float64{0, 1, 8, 27, 1000, 16581375}

	got, ok := runProcess(input)
	if !ok {
		t.Fatal("timed out waiting for pipeline results")
	}

	if len(got) != len(want) {
		t.Fatalf("got %d results, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("result[%d] = %v, want %v", i, got[i], want[i])
		}
	}
}

func TestProcessEmptyInput(t *testing.T) {
	got, ok := runProcess(nil)
	if !ok {
		t.Fatal("timed out waiting for pipeline results")
	}

	if len(got) != 0 {
		t.Fatalf("expected no results, got %d", len(got))
	}
}

func TestProcessClosesOutput(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	go process(in, out)

	close(in)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("expected output channel to be closed")
		}
	case <-time.After(time.Second):
		t.Fatal("process did not close output channel")
	}
}

func TestProcessDoesNotDeadlock(t *testing.T) {
	const count = 1000

	input := make([]uint8, count)
	for i := range input {
		input[i] = uint8(i)
	}

	got, ok := runProcess(input)
	if !ok {
		t.Fatal("process deadlocked")
	}
	if len(got) != count {
		t.Fatalf("got %d results, want %d", len(got), count)
	}
}

func runProcess(input []uint8) ([]float64, bool) {
	in := make(chan uint8)
	out := make(chan float64)
	go process(in, out)

	go func() {
		for _, n := range input {
			in <- n
		}
		close(in)
	}()

	var got []float64
	timeout := time.After(5 * time.Second)
	for {
		select {
		case res, ok := <-out:
			if !ok {
				return got, true
			}
			got = append(got, res)
		case <-timeout:
			return got, false
		}
	}
}
