package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type sequence struct {
	initialRegisters [4]int
	instruction      [4]int
	outputRegisters  [4]int
}

func (s *sequence) a() int {
	return s.instruction[1]
}

func (s *sequence) b() int {
	return s.instruction[2]
}

func (s *sequence) c() int {
	return s.instruction[3]
}

func (s *sequence) simulateOpCodes() int {
	res := 0

	// fmt.Printf("Evaluating ops for %d, %d with %v\n", s.a(), s.b(), s.initialRegisters)

	for op := opcode(0); op < nOpCodes; op++ {
		out := interpret(op, s.a(), s.b(), s.initialRegisters)

		if out == s.outputRegisters[s.c()] {
			// fmt.Printf("%2d is a match\n", op)
			res++
		}
	}

	return res
}

type opcode int

const (
	addr opcode = iota
	addi

	mulr
	muli

	banr
	bani

	borr
	bori

	setr
	seti

	gtir
	gtri
	gtrr

	eqir
	eqri
	eqrr
)

const nOpCodes = 16

func interpret(opcode opcode, A, B int, registers [4]int) int {
	switch opcode {
	case addr:
		return registers[A] + registers[B]
	case addi:
		return registers[A] + B
	case mulr:
		return registers[A] * registers[B]
	case muli:
		return registers[A] * B
	case banr:
		return registers[A] & registers[B]
	case bani:
		return registers[A] & B
	case borr:
		return registers[A] | registers[B]
	case bori:
		return registers[A] | B
	case setr:
		return registers[A]
	case seti:
		return A
	case gtir:
		if A > registers[B] {
			return 1
		}
		return 0
	case gtri:
		if registers[A] > B {
			return 1
		}
		return 0
	case gtrr:
		if registers[A] > registers[B] {
			return 1
		}
		return 0
	case eqir:
		if A == registers[B] {
			return 1
		}
		return 0
	case eqri:
		if registers[A] == B {
			return 1
		}
		return 0
	case eqrr:
		if registers[A] == registers[B] {
			return 1
		}
		return 0
	default:
		panic(fmt.Sprintf("Unknown opcode %v", opcode))
	}
}

func parseInput(r io.Reader) []sequence {
	res := make([]sequence, 0)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Before:") {
			s := sequence{
				initialRegisters: parseRegisters(line, "Before: "),
				instruction:      parseInstructions(scanner),
				outputRegisters:  parseRegisters(scanner.Text(), "After:  "),
			}
			res = append(res, s)
		}
	}

	return res
}

func parseInstructions(scanner *bufio.Scanner) [4]int {
	if !scanner.Scan() {
		panic("Expected to find a line containing the instructions")
	}

	var op, a, b, c int

	n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d", &op, &a, &b, &c)

	if n != 4 || err != nil {
		panic("Could not read instruction")
	}

	// advance it so that the next line of registers can be parsed
	if !scanner.Scan() {
		panic("Expected to find a line containing the After registers")
	}

	return [4]int{op, a, b, c}
}

func parseRegisters(line, preamble string) [4]int {
	var r0, r1, r2, r3 int

	_, err := fmt.Sscanf(line, preamble+"[%d, %d, %d, %d]", &r0, &r1, &r2, &r3)

	if err != nil {
		panic(err)
	}

	return [4]int{r0, r1, r2, r3}
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	sequences := parseInput(f)

	n := 0

	for _, s := range sequences {
		possibleOpCodes := s.simulateOpCodes()
		// fmt.Printf("%v has %d possible opcodes\n", s.instruction, possibleOpCodes)
		if possibleOpCodes >= 3 {
			n++
		}
	}

	fmt.Printf("Part 1: %v in %v\n", n, time.Since(start))
}
