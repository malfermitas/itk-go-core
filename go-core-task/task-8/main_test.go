package main

import (
	"sync"
	"testing"
	"time"
)

func TestNewCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()

	if wg == nil {
		t.Fatal("NewCustomWaitGroup returned nil")
	}
	if wg.channel == nil {
		t.Fatal("channel is not initialized")
	}
	if wg.currentDelta != 0 {
		t.Fatalf("expected currentDelta to be 0, got %d", wg.currentDelta)
	}
}

func TestAddIncrementsCounter(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(3)

	wg.mutex.Lock()
	got := wg.currentDelta
	wg.mutex.Unlock()

	if got != 3 {
		t.Fatalf("expected currentDelta to be 3, got %d", got)
	}
}

func TestDoneDecrementsCounter(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(2)

	wg.Done()

	wg.mutex.Lock()
	got := wg.currentDelta
	wg.mutex.Unlock()

	if got != 1 {
		t.Fatalf("expected currentDelta to be 1, got %d", got)
	}
}

func TestWaitReturnsAfterDone(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(1)

	finished := make(chan struct{})
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
		t.Fatal("Wait returned before Done was called")
	case <-time.After(50 * time.Millisecond):
	}

	wg.Done()

	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return after Done was called")
	}
}

func TestWaitMultipleDone(t *testing.T) {
	wg := NewCustomWaitGroup()
	const workers = 10
	wg.Add(workers)

	finished := make(chan struct{})
	go func() {
		wg.Wait()
		close(finished)
	}()

	for i := 0; i < workers-1; i++ {
		wg.Done()
		if isClosed(finished) {
			t.Fatalf("Wait returned after only %d of %d Done calls", i+1, workers)
		}
	}

	wg.Done()

	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return after all workers finished")
	}
}

func TestDonePanicsOnNegativeCounter(t *testing.T) {
	wg := NewCustomWaitGroup()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic when counter becomes negative")
		}
	}()

	wg.Done()
}

func TestConcurrentAddAndDone(t *testing.T) {
	wg := NewCustomWaitGroup()
	const workers = 100
	wg.Add(workers)

	var start sync.WaitGroup
	start.Add(1)
	for i := 0; i < workers; i++ {
		go func() {
			start.Wait()
			wg.Done()
		}()
	}

	start.Done()

	finished := make(chan struct{})
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case <-time.After(5 * time.Second):
		t.Fatal("Wait did not return after concurrent Done calls")
	}
}

func TestWaitUnblocksWhenAddThenDoneSynchronously(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(1)
	wg.Done()

	finished := make(chan struct{})
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("Wait blocked even though counter is zero")
	}
}

func isClosed(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
