package main

import (
	"fmt"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		players           int
		nMarbles          int
		expectedHighScore int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
		{424, 71482, 408679},
		{424, 7148200, 3443939356},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("with %d players and %d marbles", tt.players, tt.nMarbles), func(t *testing.T) {
			actualScore := Play(tt.players, tt.nMarbles)
			if actualScore != tt.expectedHighScore {
				t.Errorf("Expected score of %d but got %d", tt.expectedHighScore, actualScore)
			}
		})
	}
}
