package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`1, 1
		1, 6
		8, 3
		3, 4
		5, 5
		8, 9`, 17},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			area := CalculateArea(strings.NewReader(tt.input))
			if area != tt.expected {
				t.Errorf("expected %d but was %d", tt.expected, area)
			}
		})
	}
}
