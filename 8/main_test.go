package main

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	tests := []struct {
		input            string
		expectedChecksum int
		expectedValue    int
	}{
		{`2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`, 138, 66},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			analysis := Analyse(strings.NewReader(tt.input))
			if analysis.Checksum != tt.expectedChecksum {
				t.Errorf("Bad checksum: expected %d but got %d", tt.expectedChecksum, analysis.Checksum)
			}
			if analysis.Value != tt.expectedValue {
				t.Errorf("Bad value: expected %d but got %d", tt.expectedValue, analysis.Value)
			}
		})
	}
}
