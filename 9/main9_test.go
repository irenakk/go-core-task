package main

import (
	"testing"
)

func TestCubePipeline_Basic(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)

	go func() {
		input <- 2
		input <- 3
		close(input)
	}()

	expected := []float64{8, 27}
	for _, exp := range expected {
		result := <-output
		if result != exp {
			t.Errorf("got %.0f, want %.0f", result, exp)
		}
	}
}

func TestCubePipeline_ChannelClosed(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)

	close(input)

	_, ok := <-output
	if ok {
		t.Error("output channel should be closed")
	}
}
