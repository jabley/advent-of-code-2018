package main

import (
	"fmt"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		iterations int
		scoreboard string
	}{
		{
			0,
			"37",
		},
		{
			1,
			"3710",
		},
		{
			2,
			"371010",
		},
		{
			3,
			"3710101",
		},
		{
			4,
			"37101012",
		},
		{
			5,
			"371010124",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%v`, tt.iterations), func(t *testing.T) {
			p := new([]byte{'3', '7'}, 0, 1)

			for i := 0; i < tt.iterations; i++ {
				p.mix()
			}

			if tt.scoreboard != p.scoreboard() {
				t.Errorf("Expected <%s> but was <%s>", tt.scoreboard, p.scoreboard())
			}
		})
	}
}

func TestFollowingScores(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{5, "0124515891"},
		{9, "5158916779"},
		{18, "9251071085"},
		{2018, "5941429882"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%v`, tt.input), func(t *testing.T) {

			p := new([]byte{'3', '7'}, 0, 1)

			for len(p.scoreboard()) < tt.input+10 {
				p.mix()
			}

			if p.followingScores(tt.input) != tt.expected {
				t.Errorf("Expected <%s> but was <%s>", tt.expected, p.followingScores(tt.input))
			}
		})
	}
}
