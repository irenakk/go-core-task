package main

import (
	"sort"
	"testing"
	"time"
)

// helper: создаёт канал и отправляет в него значения
func makeChannel(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// helper: читает все значения из канала
func readAll(ch <-chan int) []int {
	var result []int
	for v := range ch {
		result = append(result, v)
	}
	return result
}

func TestJoinChannels_SingleChannel(t *testing.T) {
	ch := makeChannel(1, 2, 3)
	merged := joinChannels(ch)

	result := readAll(merged)
	sort.Ints(result)

	expected := []int{1, 2, 3}
	if len(result) != len(expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("got %v, want %v", result, expected)
			break
		}
	}
}

func TestJoinChannels_MultipleChannels(t *testing.T) {
	ch1 := makeChannel(1, 2, 3)
	ch2 := makeChannel(10, 20, 30)
	ch3 := makeChannel(100, 200, 300)

	merged := joinChannels(ch1, ch2, ch3)
	result := readAll(merged)
	sort.Ints(result)

	expected := []int{1, 2, 3, 10, 20, 30, 100, 200, 300}
	if len(result) != len(expected) {
		t.Errorf("got %d elements, want %d", len(result), len(expected))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("got %v, want %v", result, expected)
			break
		}
	}
}

func TestJoinChannels_NoChannels(t *testing.T) {
	merged := joinChannels()
	result := readAll(merged)

	if len(result) != 0 {
		t.Errorf("got %v, want empty slice", result)
	}
}

func TestJoinChannels_EmptyChannels(t *testing.T) {
	ch1 := makeChannel()
	ch2 := makeChannel()

	merged := joinChannels(ch1, ch2)
	result := readAll(merged)

	if len(result) != 0 {
		t.Errorf("got %v, want empty slice", result)
	}
}

func TestJoinChannels_MixedEmptyAndNonEmpty(t *testing.T) {
	ch1 := makeChannel()
	ch2 := makeChannel(1, 2, 3)
	ch3 := makeChannel()

	merged := joinChannels(ch1, ch2, ch3)
	result := readAll(merged)
	sort.Ints(result)

	expected := []int{1, 2, 3}
	if len(result) != len(expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestJoinChannels_ChannelClosed(t *testing.T) {
	ch1 := makeChannel(1)
	ch2 := makeChannel(2)

	merged := joinChannels(ch1, ch2)

	for range merged {
	}

	_, ok := <-merged
	if ok {
		t.Error("merged channel should be closed")
	}
}

func TestJoinChannels_AllValuesReceived(t *testing.T) {
	ch1 := makeChannel(1, 2)
	ch2 := makeChannel(3, 4)

	merged := joinChannels(ch1, ch2)
	result := readAll(merged)

	sum := 0
	for _, v := range result {
		sum += v
	}

	expectedSum := 1 + 2 + 3 + 4
	if sum != expectedSum {
		t.Errorf("sum = %d, want %d", sum, expectedSum)
	}
}

func TestJoinChannels_CompletesInTime(t *testing.T) {
	ch1 := makeChannel(1, 2, 3)
	ch2 := makeChannel(4, 5, 6)

	merged := joinChannels(ch1, ch2)

	done := make(chan struct{})
	go func() {
		readAll(merged)
		close(done)
	}()

	select {
	case <-done:
		// OK
	case <-time.After(2 * time.Second):
		t.Error("joinChannels did not complete in time")
	}
}

func TestJoinChannels_LargeNumberOfChannels(t *testing.T) {
	channels := make([]<-chan int, 10)
	for i := 0; i < 10; i++ {
		channels[i] = makeChannel(i * 10)
	}

	merged := joinChannels(channels...)
	result := readAll(merged)

	if len(result) != 10 {
		t.Errorf("got %d elements, want 10", len(result))
	}

	sum := 0
	for _, v := range result {
		sum += v
	}
	expectedSum := 0 + 10 + 20 + 30 + 40 + 50 + 60 + 70 + 80 + 90
	if sum != expectedSum {
		t.Errorf("sum = %d, want %d", sum, expectedSum)
	}
}

func TestJoinChannels_NegativeNumbers(t *testing.T) {
	ch1 := makeChannel(-1, -2, -3)
	ch2 := makeChannel(1, 2, 3)

	merged := joinChannels(ch1, ch2)
	result := readAll(merged)

	if len(result) != 6 {
		t.Errorf("got %d elements, want 6", len(result))
	}

	sum := 0
	for _, v := range result {
		sum += v
	}
	if sum != 0 {
		t.Errorf("sum = %d, want 0", sum)
	}
}
