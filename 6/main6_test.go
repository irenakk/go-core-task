package main

import (
	"testing"
	"time"
)

func TestRandomGenerator_Count(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{"zero elements", 0},
		{"one element", 1},
		{"ten elements", 10},
		{"hundred elements", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int)
			go RandomGenerator(ch, tt.count)

			received := 0
			for range ch {
				received++
			}

			if received != tt.count {
				t.Errorf("RandomGenerator(ch, %d): received %d elements, want %d",
					tt.count, received, tt.count)
			}
		})
	}
}

func TestRandomGenerator_ChannelClosed(t *testing.T) {
	ch := make(chan int)
	go RandomGenerator(ch, 5)

	for range ch {
	}

	_, ok := <-ch
	if ok {
		t.Error("Channel should be closed after RandomGenerator completes")
	}
}

func TestRandomGenerator_ValuesInRange(t *testing.T) {
	ch := make(chan int)
	count := 100
	go RandomGenerator(ch, count)

	for num := range ch {
		if num < 0 || num >= 100 {
			t.Errorf("RandomGenerator produced value %d, want [0, 100)", num)
		}
	}
}

func TestRandomGenerator_NonBlocking(t *testing.T) {
	ch := make(chan int)
	count := 10

	done := make(chan struct{})
	go func() {
		RandomGenerator(ch, count)
		close(done)
	}()

	received := 0
	for range ch {
		received++
	}

	select {
	case <-done:
		// OK
	case <-time.After(time.Second):
		t.Error("RandomGenerator did not complete in time")
	}

	if received != count {
		t.Errorf("Received %d elements, want %d", received, count)
	}
}

func TestRandomGenerator_UnbufferedChannel(t *testing.T) {
	ch := make(chan int)
	go RandomGenerator(ch, 5)

	count := 0
	for range ch {
		count++
	}

	if count != 5 {
		t.Errorf("Expected 5 values, got %d", count)
	}
}

func TestRandomGenerator_MultipleGoroutines(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go RandomGenerator(ch1, 5)
	go RandomGenerator(ch2, 10)

	count1, count2 := 0, 0

	done1 := make(chan struct{})
	done2 := make(chan struct{})

	go func() {
		for range ch1 {
			count1++
		}
		close(done1)
	}()

	go func() {
		for range ch2 {
			count2++
		}
		close(done2)
	}()

	<-done1
	<-done2

	if count1 != 5 {
		t.Errorf("Channel 1: got %d elements, want 5", count1)
	}
	if count2 != 10 {
		t.Errorf("Channel 2: got %d elements, want 10", count2)
	}
}
