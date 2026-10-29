package main

import (
	"testing"
	"time"
)

func TestRandomGeneratorIsUnbuffered(t *testing.T) {
	ch := RandomGenerator()
	if cap(ch) != 0 {
		t.Fatalf("expected unbuffered channel, got capacity %d", cap(ch))
	}
}

func TestRandomGeneratorValuesInRange(t *testing.T) {
	ch := RandomGenerator()
	timeout := time.After(5 * time.Second)

	for i := 0; i < 100; i++ {
		select {
		case v := <-ch:
			if v < 0 || v >= 100 {
				t.Fatalf("value %d out of range [0,100)", v)
			}
		case <-timeout:
			t.Fatal("timed out waiting for random value")
		}
	}
}

func TestTakeReturnsRequestedCount(t *testing.T) {
	ch := RandomGenerator()
	got := Take(ch, 5)
	if len(got) != 5 {
		t.Fatalf("Take returned %d values, want 5", len(got))
	}
	for _, v := range got {
		if v < 0 || v >= 100 {
			t.Fatalf("value %d out of range [0,100)", v)
		}
	}
}

func TestTakeZero(t *testing.T) {
	ch := RandomGenerator()
	got := Take(ch, 0)
	if len(got) != 0 {
		t.Fatalf("Take(0) returned %d values, want 0", len(got))
	}
}
