package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/bits"
	"os"
	"strings"
	"time"
)

// observations is a slight hack around the lack of a set type in Go.
// It is a map of opcode to the set of values that might equate to
// that opcode. Bit twiddling FTW.
type observations map[opcode]uint16

func newObservations() observations {
	res := make(observations)

	for op := addr; op < nOpCodes; op++ {
		res[op] = math.MaxUint16
	}

	return res
}

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

func (s *sequence) opcode() int {
	return s.instruction[0]
}

func (s *sequence) simulateOpCodes(observations observations) int {
	res := 0

	// fmt.Printf("Evaluating ops for %d, %d with %v\n", s.a(), s.b(), s.initialRegisters)

	for op := opcode(0); op < nOpCodes; op++ {
		bit := uint16(1 << uint16(s.opcode()))
		if observations[op]&bit != bit {
			continue
		}
		out := interpret(op, s.a(), s.b(), s.initialRegisters)

		if out == s.outputRegisters[s.c()] {
			// fmt.Printf("%2d is a match\n", op)
			res++
		} else {
			// this opcode cannot be related to this value, so clear it out
			observations[op] &^= bit
			// fmt.Printf("observations for opcode %d is now %016b after %d failed simulation\n", op, observations[op], s.opcode())
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

func parseProgram(r io.Reader) [][4]int {
	instructions := make([][4]int, 0)

	scanner := bufio.NewScanner(r)

	var op, a, b, c int

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}

		n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d", &op, &a, &b, &c)

		if n != 4 || err != nil {
			panic("Could not read instruction")
		}

		instructions = append(instructions, [4]int{op, a, b, c})
	}

	return instructions
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
	observations := newObservations()

	n := part1(sequences, observations)

	fmt.Printf("Part 1: %v in %v\n", n, time.Since(start))

	start = time.Now()

	p, err := os.Open("input_program.txt")

	if err != nil {
		panic(err)
	}

	defer p.Close()

	instructions := parseProgram(p)

	register := part2(instructions, observations)

	fmt.Printf("Part 2: %v in %v\n", register, time.Since(start))
}

func part1(sequences []sequence, observations observations) int {
	n := 0

	for _, s := range sequences {
		possibleOpCodes := s.simulateOpCodes(observations)
		// fmt.Printf("%v has %d possible opcodes\n", s.instruction, possibleOpCodes)
		if possibleOpCodes >= 3 {
			n++
		}
	}

	return n
}

func part2(instructions [][4]int, observations observations) int {
	instructionMap := decideInstructionMap(observations)

	registers := interpretInstructions([4]int{0, 0, 0, 0}, instructionMap, instructions)

	return registers[0]
}

func decideInstructionMap(observations observations) map[int]opcode {
	undecided := uint16(math.MaxUint16)

	for undecided != 0 {
		for opcode, set := range observations {
			// If the opcode operation has an unambiguous mapping to a value
			// and we haven't already removed it from the others...
			if bits.OnesCount16(set) == 1 && (undecided&set == set) {
				// fmt.Printf("%v is decided with %016b. Going to remove %016b that from all other sets\n", opcode, set, set)
				// We already have a uint16 with only a single bit set. Flip that bit in the other uint16 sets.
				undecided &^= set
				for k := range observations {
					if k == opcode {
						continue
					}
					observations[k] &^= set
				}
			}
		}
	}

	// fmt.Printf("%v\n", observations)

	instructionMap := make(map[int]opcode)

	for opcode, set := range observations {
		instructionMap[bits.Len16(set)-1] = opcode
	}

	// for i := 0; i < nOpCodes; i++ {
	// 	fmt.Printf("%d [%d]\n", i, instructionMap[i])
	// }

	return instructionMap
}

func interpretInstructions(registers [4]int, instructionMap map[int]opcode, instructions [][4]int) [4]int {
	for i := range instructions {
		// This came out nice from the part 1 implementation.
		out := interpret(instructionMap[instructions[i][0]], instructions[i][1], instructions[i][2], registers)
		registers[instructions[i][3]] = out
	}

	return registers
}
