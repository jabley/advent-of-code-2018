package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		input     string
		strongest bot
		inRange   int
	}{
		{
			`pos=<0,0,0>, r=4
		pos=<1,0,0>, r=1
		pos=<4,0,0>, r=3
		pos=<0,2,0>, r=1
		pos=<0,5,0>, r=3
		pos=<0,0,3>, r=1
		pos=<1,1,1>, r=1
		pos=<1,1,2>, r=1
		pos=<1,3,1>, r=1`,
			bot{0, 0, 0, 4},
			7,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%v`, tt.input), func(t *testing.T) {
			bots := parse(strings.NewReader(tt.input))
			strongest := findStrongest(bots)

			if strongest != tt.strongest {
				t.Fatalf("Expected %v but was %v", tt.strongest, strongest)
			}
			withinRange := withinRange(bots, strongest)

			if len(withinRange) != tt.inRange {
				t.Fatalf("Expected %v but was %v", tt.inRange, len(withinRange))
			}
		})
	}
}
