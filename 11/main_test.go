package main

import (
	"fmt"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		serial   int
		expected string
	}{
		{18, "33,45"},
		{42, "21,61"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.serial), func(t *testing.T) {
			actual := BiggestClusterFor3(tt.serial)
			if actual.String() != tt.expected {
				t.Errorf("Expected %s but was %s", tt.expected, actual)
			}
		})
	}
}
