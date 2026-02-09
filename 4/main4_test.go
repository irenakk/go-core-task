package main

import (
	"reflect"
	"testing"
)

func TestUniqueElements(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "basic case",
			slice1:   []string{"apple", "banana", "cherry", "date"},
			slice2:   []string{"banana", "date"},
			expected: []string{"apple", "cherry"},
		},
		{
			name:     "no common elements",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{"cherry", "date"},
			expected: []string{"apple", "banana"},
		},
		{
			name:     "all elements common",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{"apple", "banana"},
			expected: nil,
		},
		{
			name:     "empty first slice",
			slice1:   []string{},
			slice2:   []string{"apple", "banana"},
			expected: nil,
		},
		{
			name:     "empty second slice",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{},
			expected: []string{"apple", "banana"},
		},
		{
			name:     "both slices empty",
			slice1:   []string{},
			slice2:   []string{},
			expected: nil,
		},
		{
			name:     "slice2 has extra elements",
			slice1:   []string{"apple"},
			slice2:   []string{"apple", "banana", "cherry"},
			expected: nil,
		},
		{
			name:     "duplicates in first slice",
			slice1:   []string{"apple", "apple", "banana"},
			slice2:   []string{"banana"},
			expected: []string{"apple", "apple"},
		},
		{
			name:     "example from main",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "empty strings",
			slice1:   []string{"", "apple", ""},
			slice2:   []string{""},
			expected: []string{"apple"},
		},
		{
			name:     "single element slices - different",
			slice1:   []string{"apple"},
			slice2:   []string{"banana"},
			expected: []string{"apple"},
		},
		{
			name:     "single element slices - same",
			slice1:   []string{"apple"},
			slice2:   []string{"apple"},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniqueElements(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("uniqueElements(%v, %v) = %v, want %v",
					tt.slice1, tt.slice2, result, tt.expected)
			}
		})
	}
}

func TestUniqueElements_PreservesOrder(t *testing.T) {
	slice1 := []string{"z", "a", "m", "b"}
	slice2 := []string{"a"}
	expected := []string{"z", "m", "b"}

	result := uniqueElements(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Order not preserved: got %v, want %v", result, expected)
	}
}

func TestUniqueElements_DoesNotModifyInput(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{"banana"}

	slice1Copy := make([]string, len(slice1))
	copy(slice1Copy, slice1)
	slice2Copy := make([]string, len(slice2))
	copy(slice2Copy, slice2)

	_ = uniqueElements(slice1, slice2)

	if !reflect.DeepEqual(slice1, slice1Copy) {
		t.Error("slice1 was modified")
	}
	if !reflect.DeepEqual(slice2, slice2Copy) {
		t.Error("slice2 was modified")
	}
}
