package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			`^WNE$`,
			3,
		},
		{
			`^ENWWW(NEEE|SSE(EE|N))$`,
			10,
		},
		{
			`^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$`,
			18,
		},
		{
			`^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$`,
			23,
		},
		{
			`^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$`,
			31,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%v`, tt.input), func(t *testing.T) {
			output := path(strings.NewReader(tt.input))

			if max(output) != tt.expected {
				t.Errorf("Expected %d but was %d for %q", tt.expected, max(output), tt.input)
			}
		})
	}
}
