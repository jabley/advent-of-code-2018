package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestOpCodes(t *testing.T) {
	tests := []struct {
		initialRegisters [4]int
		opcode           opcode
		A, B             int
		output           int
	}{
		{
			[4]int{3, 2, 1, 0},
			addr,
			1,
			2,
			3,
		},
		{
			[4]int{3, 2, 1, 0},
			addi,
			1,
			2,
			4,
		},
		{
			[4]int{3, 2, 1, 0},
			mulr,
			1,
			2,
			2,
		},
		{
			[4]int{3, 2, 1, 0},
			muli,
			1,
			2,
			4,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%v`, tt.initialRegisters), func(t *testing.T) {
			output := interpret(tt.opcode, tt.A, tt.B, tt.initialRegisters)

			if output != tt.output {
				t.Errorf("Expected %v when applying %v to %v but got %v", tt.output, tt.opcode, tt.initialRegisters, output)
			}
		})
	}
}

func TestPart1Examples(t *testing.T) {
	tests := []struct {
		input           string
		possibleOpCodes int
	}{
		{
			`
Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]
`,
			3,
		},
		{`
Before: [2, 3, 1, 3]
9 2 1 1
After:  [2, 1, 1, 3]
`,
			6,
		},
		{
			`
Before: [0, 1, 0, 2]
14 3 1 0
After:  [3, 1, 0, 2]
			`,
			5,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%v`, tt.input), func(t *testing.T) {
			sequences := parseInput(strings.NewReader(tt.input))

			if len(sequences) != 1 {
				t.Errorf("Expected a single sequence but got %d", len(sequences))
			}

			possibleOpCodes := sequences[0].simulateOpCodes()

			if possibleOpCodes != tt.possibleOpCodes {
				t.Errorf("Expected %v but got %v", tt.possibleOpCodes, possibleOpCodes)
			}
		})
	}
}
