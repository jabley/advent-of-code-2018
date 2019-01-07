package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type mode int

const (
	normal mode = iota
	part1
	part2
)

type instruction struct {
	desc   string
	opcode opcode
}

func (i instruction) String() string {
	return i.desc
}

type opcode func(before []int) []int

type program struct {
	ipReg              int // the register which is bound to the instruction pointer
	instructionPointer int // the value of the instruction pointer
	instructions       []instruction
	registers          []int

	// instruction 28 compares r0 and r1 for equality. We control r0 as part of our
	// attack, so we need to track which values we see in r1 for part 2.
	seenR1Values map[int]struct{}
	lastR1Value  int
}

func (p *program) step(mode mode) (bool, int) {
	// When the instruction pointer is bound to a register, its value is
	// written to that register just before each instruction is executed
	p.registers[p.ipReg] = p.instructionPointer

	// From analysing the input program, we need to compare r1 with r0 using eqrr
	if p.instructionPointer == 28 {
		if mode == part1 {
			// So we interpret the program until we hit that instruction, then
			// dump the register value
			return false, p.registers[1]
		} else if mode == part2 {
			// For part 2, we keep track of the seen values, and note when the
			// values start to repeat. The first value which repeats is the lowest
			// non-negative value that we can put in r0 as part of the attack.
			if _, ok := p.seenR1Values[p.registers[1]]; ok {
				// fmt.Printf("repeating: %d\n", p.lastR1Value)
				return false, p.lastR1Value
			}
			p.seenR1Values[p.registers[1]] = struct{}{}
			p.lastR1Value = p.registers[1]
		}
		// fmt.Printf("%v\n", p.registers[1])
	}

	if p.instructionPointer == 17 {
		// We have a hot loop in the program
		//
		// r2 = 0
		// while ((r2+1) * 256) < r5:
		// 	r2 = r2 + 1
		// r5 = r2
		//
		// This can can be more simply written as:
		//
		// r5 = r5 / 256
		//
		// This change solves part 1 in 35 instructions rather than
		// 1848 instructions.
		//
		// It has a similar effect on solving part 2. Part 2 runtime
		// has gone from ~20 seconds to 4ms.
		// This is equivalent to inline assembly for our ElfCode,
		// rewriting the hot loop in Go. We could also have added a
		// new opcode to the language, so support native division. But
		// would have required more effort to preprocess the input.
		p.registers[5] = p.registers[5] / 256
		p.instructionPointer = 27
	}

	// if p.instructionPointer == 27 {
	// 	// Ensure that our rewrite of the hot loop hasn't mangled the registers
	// 	fmt.Printf("Leaving hot loop: %v\n", p.registers)
	// }

	instruction := p.instructions[p.instructionPointer]

	// fmt.Printf("ip=%d %v %s ", p.instructionPointer, p.registers, instruction)
	registers := instruction.opcode(p.registers)
	p.registers = registers
	// fmt.Printf("%v\n", p.registers)

	// When the instruction pointer is bound to a register, ... the value
	// of that register is written back to the instruction pointer
	// immediately after each instruction finishes execution
	p.instructionPointer = p.registers[p.ipReg]

	// Afterward, move to the next instruction by adding one to the
	// instruction pointer, even if the value in the instruction pointer
	// was just updated by an instruction.
	p.instructionPointer++

	// If the instruction pointer ever causes the device to attempt to
	// load an instruction outside the instructions defined in the
	// program, the program instead immediately halts.
	if p.instructionPointer > len(p.instructions)-1 {
		return false, -1
	}

	return true, -1
}

func (p *program) execute(mode mode) (int, int, bool) {
	if p.registers == nil {
		p.registers = []int{0, 0, 0, 0, 0, 0}
	}

	i := 1
	running := true
	res := 0
	for ; running && i < 10000000000; i++ {
		running, res = p.step(mode)
	}

	return i, res, !running
}

func parseInput(r io.Reader) *program {
	res := program{
		seenR1Values: make(map[int]struct{}),
	}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			res.ipReg = atoi(scanner.Text()[4:5])
			continue
		}

		fields := strings.Split(scanner.Text(), " ")
		res.instructions = append(res.instructions,
			parseInstruction(fields[0], atoi(fields[1]), atoi(fields[2]), atoi(fields[3])))
	}

	return &res
}

func parseInstruction(name string, a, b, c int) instruction {
	return instruction{
		desc:   fmt.Sprintf("%s %d %d %d", name, a, b, c),
		opcode: createOpcode(name, a, b, c),
	}
}

func createOpcode(name string, a, b, c int) opcode {
	switch name {
	case "addr":
		return func(before []int) []int {
			res := before
			res[c] = before[a] + before[b]
			return res
		}
	case "addi":
		return func(before []int) []int {
			res := before
			res[c] = before[a] + b
			return res
		}
	case "mulr":
		return func(before []int) []int {
			res := before
			res[c] = before[a] * before[b]
			return res
		}
	case "muli":
		return func(before []int) []int {
			res := before
			res[c] = before[a] * b
			return res
		}
	case "banr":
		return func(before []int) []int {
			res := before
			res[c] = before[a] & before[b]
			return res
		}
	case "bani":
		return func(before []int) []int {
			res := before
			res[c] = before[a] & b
			return res
		}
	case "borr":
		return func(before []int) []int {
			res := before
			res[c] = before[a] | before[b]
			return res
		}
	case "bori":
		return func(before []int) []int {
			res := before
			res[c] = before[a] | b
			return res
		}
	case "setr":
		return func(before []int) []int {
			res := before
			res[c] = before[a]
			return res
		}
	case "seti":
		return func(before []int) []int {
			res := before
			res[c] = a
			return res
		}
	case "gtir":
		return func(before []int) []int {
			res := before
			if a > before[b] {
				res[c] = 1
			} else {
				res[c] = 0
			}
			return res
		}
	case "gtri":
		return func(before []int) []int {
			res := before
			if before[a] > b {
				res[c] = 1
			} else {
				res[c] = 0
			}
			return res
		}
	case "gtrr":
		return func(before []int) []int {
			res := before
			if before[a] > before[b] {
				res[c] = 1
			} else {
				res[c] = 0
			}
			return res
		}
	case "eqir":
		return func(before []int) []int {
			res := before
			if a == before[b] {
				res[c] = 1
			} else {
				res[c] = 0
			}
			return res
		}
	case "eqri":
		return func(before []int) []int {
			res := before
			if before[a] == b {
				res[c] = 1
			} else {
				res[c] = 0
			}
			return res
		}
	case "eqrr":
		return func(before []int) []int {
			res := before
			if before[a] == before[b] {
				res[c] = 1
			} else {
				res[c] = 0
			}
			return res
		}
	default:
		panic(fmt.Sprintf("Unknown opcode %v", name))
	}
}

func atoi(a string) int {
	n, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	program := parseInput(f)

	program.registers = []int{0, 0, 0, 0, 0, 0}
	instructionCount, reg1, _ := program.execute(part1)

	fmt.Printf("Part 1 in %v: %d in %d instructions\n", time.Since(start), reg1, instructionCount)

	start = time.Now()
	program.registers = []int{0, 0, 0, 0, 0, 0}
	program.instructionPointer = 0
	instructionCount, reg1, _ = program.execute(part2)

	fmt.Printf("Part 2 in %v: %d in %d instructions\n", time.Since(start), reg1, instructionCount)
}
