package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestHashRuneSliceWithSalt(t *testing.T) {
	tests := []struct {
		name     string
		runes    []rune
		salt     string
		expected string
	}{
		{
			name:  "basic test with go-2024 salt",
			runes: []rune("Hello"),
			salt:  "go-2024",
			// "He" + "go-2024" + "llo" = "Hego-2024llo"
			expected: computeExpectedHash("Hego-2024llo"),
		},
		{
			name:  "empty rune slice",
			runes: []rune(""),
			salt:  "go-2024",
			// "" + "go-2024" + "" = "go-2024"
			expected: computeExpectedHash("go-2024"),
		},
		{
			name:  "single character",
			runes: []rune("A"),
			salt:  "go-2024",
			// "" + "go-2024" + "A" = "go-2024A"
			expected: computeExpectedHash("go-2024A"),
		},
		{
			name:  "even length string",
			runes: []rune("ABCD"),
			salt:  "go-2024",
			// "AB" + "go-2024" + "CD" = "ABgo-2024CD"
			expected: computeExpectedHash("ABgo-2024CD"),
		},
		{
			name:  "with cyrillic characters",
			runes: []rune("Привет"),
			salt:  "go-2024",
			// "При" + "go-2024" + "вет" = "Приgo-2024вет"
			expected: computeExpectedHash("Приgo-2024вет"),
		},
		{
			name:  "empty salt",
			runes: []rune("Test"),
			salt:  "",
			// "Te" + "" + "st" = "Test"
			expected: computeExpectedHash("Test"),
		},
		{
			name:  "real data from main",
			runes: []rune("4242423.14Golangtrue(1+2i)"),
			salt:  "go-2024",
			// 26 chars total, mid=13, first 13 chars "4242423.14Gol" + salt + "angtrue(1+2i)"
			expected: computeExpectedHash("4242423.14Golgo-2024angtrue(1+2i)"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HashRuneSliceWithSalt(tt.runes, tt.salt)
			if result != tt.expected {
				t.Errorf("HashRuneSliceWithSalt(%q, %q) = %v, want %v",
					string(tt.runes), tt.salt, result, tt.expected)
			}
		})
	}
}

func TestHashRuneSliceWithSalt_HashLength(t *testing.T) {
	result := HashRuneSliceWithSalt([]rune("test"), "salt")
	// SHA256 produces 32 bytes = 64 hex characters
	if len(result) != 64 {
		t.Errorf("Expected hash length 64, got %d", len(result))
	}
}

func TestHashRuneSliceWithSalt_Deterministic(t *testing.T) {
	runes := []rune("TestString")
	salt := "go-2024"

	result1 := HashRuneSliceWithSalt(runes, salt)
	result2 := HashRuneSliceWithSalt(runes, salt)

	if result1 != result2 {
		t.Error("Hash function is not deterministic")
	}
}

func TestHashRuneSliceWithSalt_DifferentInputs(t *testing.T) {
	salt := "go-2024"

	result1 := HashRuneSliceWithSalt([]rune("Input1"), salt)
	result2 := HashRuneSliceWithSalt([]rune("Input2"), salt)

	if result1 == result2 {
		t.Error("Different inputs produced the same hash")
	}
}

func TestHashRuneSliceWithSalt_DifferentSalts(t *testing.T) {
	runes := []rune("SameInput")

	result1 := HashRuneSliceWithSalt(runes, "salt1")
	result2 := HashRuneSliceWithSalt(runes, "salt2")

	if result1 == result2 {
		t.Error("Different salts produced the same hash")
	}
}

// computeExpectedHash вспомогательная функция для вычисления ожидаемого хэша
func computeExpectedHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
