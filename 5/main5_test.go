package main

import (
	"reflect"
	"testing"
)

func TestUniqueElements(t *testing.T) {
	tests := []struct {
		name             string
		slice1           []int
		slice2           []int
		expectedBool     bool
		expectedIntersec []int
	}{
		{
			name:             "example from task",
			slice1:           []int{65, 3, 58, 678, 64},
			slice2:           []int{64, 2, 3, 43},
			expectedBool:     true,
			expectedIntersec: []int{3, 64},
		},
		{
			name:             "no intersection",
			slice1:           []int{1, 2, 3},
			slice2:           []int{4, 5, 6},
			expectedBool:     false,
			expectedIntersec: nil,
		},
		{
			name:             "all elements intersect",
			slice1:           []int{1, 2, 3},
			slice2:           []int{1, 2, 3},
			expectedBool:     true,
			expectedIntersec: []int{1, 2, 3},
		},
		{
			name:             "empty first slice",
			slice1:           []int{},
			slice2:           []int{1, 2, 3},
			expectedBool:     false,
			expectedIntersec: nil,
		},
		{
			name:             "empty second slice",
			slice1:           []int{1, 2, 3},
			slice2:           []int{},
			expectedBool:     false,
			expectedIntersec: nil,
		},
		{
			name:             "both slices empty",
			slice1:           []int{},
			slice2:           []int{},
			expectedBool:     false,
			expectedIntersec: nil,
		},
		{
			name:             "single common element",
			slice1:           []int{1, 2, 3, 4, 5},
			slice2:           []int{3},
			expectedBool:     true,
			expectedIntersec: []int{3},
		},
		{
			name:             "duplicates in first slice",
			slice1:           []int{1, 1, 2, 2, 3},
			slice2:           []int{1, 3},
			expectedBool:     true,
			expectedIntersec: []int{1, 1, 3},
		},
		{
			name:             "negative numbers",
			slice1:           []int{-5, -3, 0, 3, 5},
			slice2:           []int{-3, 0, 10},
			expectedBool:     true,
			expectedIntersec: []int{-3, 0},
		},
		{
			name:             "single element slices - same",
			slice1:           []int{42},
			slice2:           []int{42},
			expectedBool:     true,
			expectedIntersec: []int{42},
		},
		{
			name:             "single element slices - different",
			slice1:           []int{42},
			slice2:           []int{24},
			expectedBool:     false,
			expectedIntersec: nil,
		},
		{
			name:             "zeros",
			slice1:           []int{0, 1, 2},
			slice2:           []int{0, 3, 4},
			expectedBool:     true,
			expectedIntersec: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotSlice := uniqueElements(tt.slice1, tt.slice2)

			if gotBool != tt.expectedBool {
				t.Errorf("uniqueElements(%v, %v) bool = %v, want %v",
					tt.slice1, tt.slice2, gotBool, tt.expectedBool)
			}

			if !reflect.DeepEqual(gotSlice, tt.expectedIntersec) {
				t.Errorf("uniqueElements(%v, %v) slice = %v, want %v",
					tt.slice1, tt.slice2, gotSlice, tt.expectedIntersec)
			}
		})
	}
}

func TestUniqueElements_PreservesOrder(t *testing.T) {
	slice1 := []int{5, 3, 1, 4, 2}
	slice2 := []int{1, 2, 3}
	expectedSlice := []int{3, 1, 2}

	_, result := uniqueElements(slice1, slice2)

	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Order not preserved: got %v, want %v", result, expectedSlice)
	}
}

func TestUniqueElements_DoesNotModifyInput(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4}

	slice1Copy := make([]int, len(slice1))
	copy(slice1Copy, slice1)
	slice2Copy := make([]int, len(slice2))
	copy(slice2Copy, slice2)

	_, _ = uniqueElements(slice1, slice2)

	if !reflect.DeepEqual(slice1, slice1Copy) {
		t.Error("slice1 was modified")
	}
	if !reflect.DeepEqual(slice2, slice2Copy) {
		t.Error("slice2 was modified")
	}
}

func TestUniqueElements_BoolConsistency(t *testing.T) {
	// Проверяем что bool=true только когда слайс не пустой
	testCases := []struct {
		slice1 []int
		slice2 []int
	}{
		{[]int{1, 2, 3}, []int{2}},
		{[]int{1, 2, 3}, []int{4, 5, 6}},
		{[]int{}, []int{1, 2}},
	}

	for _, tc := range testCases {
		gotBool, gotSlice := uniqueElements(tc.slice1, tc.slice2)

		expectedBool := len(gotSlice) > 0
		if gotBool != expectedBool {
			t.Errorf("uniqueElements(%v, %v): bool=%v but slice=%v (len=%d)",
				tc.slice1, tc.slice2, gotBool, gotSlice, len(gotSlice))
		}
	}
}
