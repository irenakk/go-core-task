package main

import (
	"reflect"
	"testing"
)

func newStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    int
		expected map[string]int
	}{
		{
			name:     "add single element",
			key:      "apple",
			value:    5,
			expected: map[string]int{"apple": 5},
		},
		{
			name:     "add with zero value",
			key:      "zero",
			value:    0,
			expected: map[string]int{"zero": 0},
		},
		{
			name:     "add with negative value",
			key:      "negative",
			value:    -10,
			expected: map[string]int{"negative": -10},
		},
		{
			name:     "add with empty key",
			key:      "",
			value:    42,
			expected: map[string]int{"": 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newStringIntMap()
			s.Add(tt.key, tt.value)

			if s.data[tt.key] != tt.value {
				t.Errorf("Add(%q, %d): got %d, want %d", tt.key, tt.value, s.data[tt.key], tt.value)
			}
		})
	}
}

func TestAdd_MultipleElements(t *testing.T) {
	s := newStringIntMap()
	s.Add("apple", 5)
	s.Add("banana", 3)
	s.Add("cherry", 7)

	expected := map[string]int{"apple": 5, "banana": 3, "cherry": 7}

	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("After adding multiple elements: got %v, want %v", s.data, expected)
	}
}

func TestAdd_OverwriteExisting(t *testing.T) {
	s := newStringIntMap()
	s.Add("key", 10)
	s.Add("key", 20)

	if s.data["key"] != 20 {
		t.Errorf("Add should overwrite existing key: got %d, want 20", s.data["key"])
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name      string
		initial   map[string]int
		removeKey string
		expected  map[string]int
	}{
		{
			name:      "remove existing key",
			initial:   map[string]int{"apple": 5, "banana": 3},
			removeKey: "apple",
			expected:  map[string]int{"banana": 3},
		},
		{
			name:      "remove non-existing key",
			initial:   map[string]int{"apple": 5},
			removeKey: "banana",
			expected:  map[string]int{"apple": 5},
		},
		{
			name:      "remove from empty map",
			initial:   map[string]int{},
			removeKey: "apple",
			expected:  map[string]int{},
		},
		{
			name:      "remove last element",
			initial:   map[string]int{"apple": 5},
			removeKey: "apple",
			expected:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{data: tt.initial}
			s.Remove(tt.removeKey)

			if !reflect.DeepEqual(s.data, tt.expected) {
				t.Errorf("Remove(%q): got %v, want %v", tt.removeKey, s.data, tt.expected)
			}
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[string]int
		checkKey string
		expected bool
	}{
		{
			name:     "key exists",
			initial:  map[string]int{"apple": 5},
			checkKey: "apple",
			expected: true,
		},
		{
			name:     "key does not exist",
			initial:  map[string]int{"apple": 5},
			checkKey: "banana",
			expected: false,
		},
		{
			name:     "empty map",
			initial:  map[string]int{},
			checkKey: "apple",
			expected: false,
		},
		{
			name:     "key with zero value exists",
			initial:  map[string]int{"zero": 0},
			checkKey: "zero",
			expected: true,
		},
		{
			name:     "empty string key exists",
			initial:  map[string]int{"": 42},
			checkKey: "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{data: tt.initial}
			result := s.Exists(tt.checkKey)

			if result != tt.expected {
				t.Errorf("Exists(%q): got %v, want %v", tt.checkKey, result, tt.expected)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name          string
		initial       map[string]int
		getKey        string
		expectedValue int
		expectedOk    bool
	}{
		{
			name:          "get existing key",
			initial:       map[string]int{"apple": 5},
			getKey:        "apple",
			expectedValue: 5,
			expectedOk:    true,
		},
		{
			name:          "get non-existing key",
			initial:       map[string]int{"apple": 5},
			getKey:        "banana",
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			name:          "get from empty map",
			initial:       map[string]int{},
			getKey:        "apple",
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			name:          "get key with zero value",
			initial:       map[string]int{"zero": 0},
			getKey:        "zero",
			expectedValue: 0,
			expectedOk:    true,
		},
		{
			name:          "get key with negative value",
			initial:       map[string]int{"negative": -100},
			getKey:        "negative",
			expectedValue: -100,
			expectedOk:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{data: tt.initial}
			value, ok := s.Get(tt.getKey)

			if value != tt.expectedValue || ok != tt.expectedOk {
				t.Errorf("Get(%q): got (%d, %v), want (%d, %v)",
					tt.getKey, value, ok, tt.expectedValue, tt.expectedOk)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[string]int
		expected map[string]int
	}{
		{
			name:     "copy non-empty map",
			initial:  map[string]int{"apple": 5, "banana": 3},
			expected: map[string]int{"apple": 5, "banana": 3},
		},
		{
			name:     "copy empty map",
			initial:  map[string]int{},
			expected: map[string]int{},
		},
		{
			name:     "copy single element",
			initial:  map[string]int{"only": 1},
			expected: map[string]int{"only": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIntMap{data: tt.initial}
			result := s.Copy()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Copy(): got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCopy_Independence(t *testing.T) {
	s := newStringIntMap()
	s.Add("apple", 5)
	s.Add("banana", 3)

	copied := s.Copy()

	s.Add("cherry", 7)
	s.Remove("apple")

	if _, exists := copied["cherry"]; exists {
		t.Error("Copy was affected by adding to original")
	}

	if _, exists := copied["apple"]; !exists {
		t.Error("Copy was affected by removing from original")
	}
}
