package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestSema_WaitForGoroutines(t *testing.T) {
	sem := New(5)
	var counter int32

	for i := 0; i < 5; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			sem.Inc(1)
		}()
	}

	sem.Dec(5)

	if counter != 5 {
		t.Errorf("counter = %d, want 5", counter)
	}
}

func TestSema_IncDec(t *testing.T) {
	sem := New(3)

	done := make(chan struct{})
	go func() {
		sem.Inc(3)
		close(done)
	}()

	select {
	case <-done:
		// OK
	case <-time.After(time.Second):
		t.Error("Inc should not block when buffer has space")
	}

	sem.Dec(3)

	// Канал должен быть пустым
	select {
	case <-sem:
		t.Error("Semaphore should be empty after Dec")
	default:
		// OK
	}
}
