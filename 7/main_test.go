package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	tests := []struct {
		input           string
		expectedOrder   string
		workerCount     int
		fixedJobCost    int
		expectedSeconds int
	}{
		{
			`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`,
			"CABDFE", 2, 0, 15},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			order := CalculateOrder(strings.NewReader(tt.input))
			if order != tt.expectedOrder {
				t.Errorf("Expected <%s> but got <%s>", tt.expectedOrder, order)
			}
			duration := CalculateTime(strings.NewReader(tt.input), tt.workerCount, tt.fixedJobCost)
			if duration != tt.expectedSeconds {
				t.Errorf("Expected <%d> but got <%d>", tt.expectedSeconds, duration)
			}
		})
	}
}
