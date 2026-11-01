package main

import (
	"reflect"
	"testing"
)

func TestRandSlice(t *testing.T) {
	length := 10
	result := randSlice(length)

	if len(result) != length {
		t.Errorf("randSlice(%d) returned slice of length %d, want %d", length, len(result), length)
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "mixed numbers",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "all even",
			input:    []int{2, 4, 6, 8},
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "all odd",
			input:    []int{1, 3, 5, 7},
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "negative even numbers",
			input:    []int{-4, -3, -2, -1, 0},
			expected: []int{-4, -2, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceExample(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sliceExample(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		num      int
		expected []int
	}{
		{
			name:     "add to non-empty slice",
			input:    []int{1, 2, 3},
			num:      4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "add to empty slice",
			input:    []int{},
			num:      1,
			expected: []int{1},
		},
		{
			name:     "add negative number",
			input:    []int{1, 2},
			num:      -5,
			expected: []int{1, 2, -5},
		},
		{
			name:     "add zero",
			input:    []int{1, 2, 3},
			num:      0,
			expected: []int{1, 2, 3, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addElements(tt.input, tt.num)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("addElements(%v, %d) = %v, want %v", tt.input, tt.num, result, tt.expected)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "non-empty slice",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "empty slice",
			input: []int{},
		},
		{
			name:  "single element",
			input: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copySlice(tt.input)
			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("copySlice(%v) = %v, want %v", tt.input, result, tt.input)
			}
		})
	}
}

func TestCopySlice_Independence(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	copied := copySlice(original)

	original[0] = 999

	if copied[0] == 999 {
		t.Error("Copy was affected by changes to original slice")
	}

	if copied[0] != 1 {
		t.Errorf("copied[0] = %d, want 1", copied[0])
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{
			name:     "remove from middle",
			input:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "remove first element",
			input:    []int{1, 2, 3, 4, 5},
			index:    0,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "remove last element",
			input:    []int{1, 2, 3, 4, 5},
			index:    4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "single element slice",
			input:    []int{42},
			index:    0,
			expected: []int{},
		},
		{
			name:     "negative index - return original",
			input:    []int{1, 2, 3},
			index:    -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "index out of bounds - return original",
			input:    []int{1, 2, 3},
			index:    10,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElement(tt.input, tt.index)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("removeElement(%v, %d) = %v, want %v", tt.input, tt.index, result, tt.expected)
			}
		})
	}
}

func TestRemoveElement_DoesNotModifyOriginal(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	_ = removeElement(original, 2)

	if !reflect.DeepEqual(original, originalCopy) {
		t.Error("removeElement modified the original slice")
	}
}
