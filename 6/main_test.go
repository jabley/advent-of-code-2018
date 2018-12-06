package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	tests := []struct {
		input              string
		maxDistance        int
		expectedArea       int
		expectedRegionSize int
	}{
		{`1, 1
		1, 6
		8, 3
		3, 4
		5, 5
		8, 9`, 32, 17, 16},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			area, regionSize := CalculateArea(strings.NewReader(tt.input), tt.maxDistance)
			if area != tt.expectedArea {
				t.Errorf("expected %d but was %d", tt.expectedArea, area)
			}
			if regionSize != tt.expectedRegionSize {
				t.Errorf("expected %d but was %d", tt.expectedRegionSize, regionSize)
			}
		})
	}
}
