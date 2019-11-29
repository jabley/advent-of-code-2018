package main

import (
	"fmt"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		depth             int
		targetX, targetY  int
		expectedRiskLevel int
		shortestTime      int
	}{
		{
			510,
			10, 10,
			114,
			45,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf(`%v`, i), func(t *testing.T) {
			constraints := &input{
				depth: tt.depth,
				x:     tt.targetX,
				y:     tt.targetY,
			}

			tiles := buildMap(constraints)

			actualRiskLevel := calculateRiskLevel(tiles, constraints)

			if actualRiskLevel != tt.expectedRiskLevel {
				t.Errorf("Expected %d but was %d for %v", tt.expectedRiskLevel, actualRiskLevel, constraints)
			}

			shortestTime := calculateShortestTime(tiles, constraints)

			if shortestTime != tt.shortestTime {
				t.Errorf("Expected %d but was %d for %v", tt.shortestTime, shortestTime, constraints)
			}
		})
	}
}
