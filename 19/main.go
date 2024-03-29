package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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
}

func (p *program) step(optimisedVersion bool) bool {
	if p.instructionPointer == 1 && optimisedVersion {
		// about to spend ages factoring r3 (from my puzzle input), so
		// just do the idiomatic go version, rather than trying to
		// execute the inefficient VM version.
		p.registers[0] = calculateSumOfFactorsOf(p.registers[3])
		return false
	}
	// When the instruction pointer is bound to a register, its value is
	// written to that register just before each instruction is executed
	p.registers[p.ipReg] = p.instructionPointer

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
		return false
	}

	return true
}

func (p *program) execute(optimisedVersion bool) int {
	if p.registers == nil {
		p.registers = []int{0, 0, 0, 0, 0, 0}
	}

	for p.step(optimisedVersion) {

	}

	return p.registers[0]
}

func parseInput(r io.Reader) *program {
	res := program{}

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

func calculateSumOfFactorsOf(nr int) int {
	// fmt.Printf("Calculating sum of divisors for %d\n", nr)
	sqrt := int(math.Sqrt(float64(nr)))
	result := 0

	for i := 1; i <= sqrt; i++ {
		if nr%i == 0 {
			result += i + nr/i
		}
	}

	return result
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	program := parseInput(f)

	output := program.execute(false)

	fmt.Printf("Part 1 in %v: %d\n", time.Since(start), output)

	start = time.Now()
	program.registers = []int{1, 0, 0, 0, 0, 0}
	program.instructionPointer = program.registers[program.ipReg]
	output = program.execute(true)
	fmt.Printf("Part 2 in %v: %d\n", time.Since(start), output)
}
