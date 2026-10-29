package main

import (
	"fmt"
	"sync"
	"time"
)

type CustomWaitGroup struct {
	channel      chan struct{}
	currentDelta int
	mutex        sync.Mutex
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{channel: make(chan struct{}), currentDelta: 0}
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.mutex.Lock()
	if wg.currentDelta == 0 && delta > 0 {
		wg.channel = make(chan struct{})
	}
	wg.currentDelta += delta
	if wg.currentDelta < 0 {
		wg.mutex.Unlock()
		panic("negative WaitGroup counter")
	}
	wg.mutex.Unlock()
}

func (wg *CustomWaitGroup) Done() {
	wg.mutex.Lock()
	wg.currentDelta--
	if wg.currentDelta < 0 {
		wg.mutex.Unlock()
		panic("negative WaitGroup counter")
	} else if wg.currentDelta == 0 {
		close(wg.channel)
	}
	wg.mutex.Unlock()
}

func (wg *CustomWaitGroup) Wait() {
	wg.mutex.Lock()
	if wg.currentDelta == 0 {
		wg.mutex.Unlock()
		return
	}
	channel := wg.channel
	wg.mutex.Unlock()
	<-channel
}

func main() {
	wg := NewCustomWaitGroup()
	wg.Add(2)
	res1, res2 := 0, 0
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		for i := range 10 {
			res1 += i * 2
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(4 * time.Second)
		for i := range 20 {
			res2 += i * 2
		}
	}()
	wg.Wait()
	fmt.Println(res1, res2)
}
