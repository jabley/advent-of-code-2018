package main

import (
	"strings"
	"testing"
)

func TestBackgroundProcess(t *testing.T) {
	tests := []struct {
		input         string
		ipReg         int
		nInstructions int
		output        int
	}{
		{
			`#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`,
			0,
			7,
			6,
		},
		{
			`#ip 2
addi 2 16 2
seti 1 2 4
seti 1 8 1
mulr 4 1 5
eqrr 5 3 5
addr 5 2 2
addi 2 1 2
addr 4 0 0
addi 1 1 1
gtrr 1 3 5
addr 2 5 2
seti 2 6 2
addi 4 1 4
gtrr 4 3 5
addr 5 2 2
seti 1 2 2
mulr 2 2 2
addi 3 2 3
mulr 3 3 3
mulr 2 3 3
muli 3 11 3
addi 5 2 5
mulr 5 2 5
addi 5 8 5
addr 3 5 3
addr 2 0 2
seti 0 4 2
setr 2 5 5
mulr 5 2 5
addr 2 5 5
mulr 2 5 5
muli 5 14 5
mulr 5 2 5
addr 3 5 3
seti 0 8 0
seti 0 5 2`,
			2,
			36,
			2280,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			program := parseInput(strings.NewReader(tt.input))

			if len(program.instructions) != tt.nInstructions {
				t.Errorf("expected %d instructions but was %d", tt.nInstructions, len(program.instructions))
			}

			if program.ipReg != tt.ipReg {
				t.Errorf("expected ip reg of %d but was %d", tt.ipReg, program.instructionPointer)
			}

			output := program.execute()

			if output != tt.output {
				t.Errorf("Expected %d but was %d", tt.output, output)
			}
		})
	}
}
