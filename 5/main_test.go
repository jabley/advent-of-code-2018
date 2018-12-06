package main

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
		{"aA", ""},
		{"abBA", ""},
		{"abAB", "abAB"},
		{"aabAAB", "aabAAB"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			stable := React(tt.input)

			if stable != tt.expected {
				t.Errorf("Expected <%s> but got <%s>", tt.expected, stable)
			}
		})
	}
}
