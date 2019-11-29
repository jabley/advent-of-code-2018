// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	desc    string
	a, b, c int
}

type program struct {
	ipReg        int // the register which is bound to the instruction pointer
	instructions []instruction
}

func (p *program) generateVanillaCode() ([]string, map[int]struct{}) {
	lines := make([]string, 0)
	// Keep track of which registers are actually used
	registers := make(map[int]struct{})

	for _, i := range p.instructions {
		if _, ok := registers[i.c]; !ok {
			registers[i.c] = struct{}{}
		}
		lines = append(lines, generateCode(p, i))
	}

	return lines, registers
}

func (p *program) loopPass(lines []string) {
	for j, i := range p.instructions {
		if i.desc == "seti" && i.c == p.ipReg && i.a < j {
			lines[i.a] += "\nfor {"
			lines[j] = "}"
		}

		// Check we have a run of 3 instructions to compare
		if j < 1 || j >= len(p.instructions)-1 {
			continue
		}

		previous := p.instructions[j-1]
		next := p.instructions[j+1]
		if previous.desc == "eqri" || previous.desc == "eqrr" {
			if i.desc == "addr" && i.c == p.ipReg {
				var other int
				if i.a == p.ipReg {
					other = i.b
				} else {
					other = i.a
				}

				if next.desc == "seti" && next.c == p.ipReg {
					if previous.c == other && next.a < j {
						lines[j] = fmt.Sprintf("if r%d == 1 { break }", other)
					}
				}
			}

		}
	}
}

func (p *program) transpile() []byte {
	var buf bytes.Buffer

	lines, registers := p.generateVanillaCode()

	p.loopPass(lines)

	buf.WriteString("// +build ignore\n\n")
	buf.WriteString("package main\n\n")
	buf.WriteString("func main() {\n")
	buf.WriteString("\tvar r0 int\n")

	// Declare the registers which are used
	for k := range registers {
		buf.WriteString(fmt.Sprintf("\tvar r%d int\n", k))
	}

	buf.WriteString("\n")

	for _, line := range lines {
		buf.WriteString("\t" + line)
		buf.WriteString("\n")
	}

	buf.WriteString("}\n")

	return buf.Bytes()
}

func generateCode(p *program, i instruction) string {
	switch i.desc {
	case "addi":
		return fmt.Sprintf("r%d = r%d + %d", i.c, i.a, i.b)
	case "addr":
		return fmt.Sprintf("r%d = r%d + r%d", i.c, i.a, i.b)
	case "bani":
		return fmt.Sprintf("r%d = r%d & %d", i.c, i.a, i.b)
	case "bori":
		return fmt.Sprintf("r%d = r%d | %d", i.c, i.a, i.b)
	case "eqri":
		return fmt.Sprintf("if r%d == %d {\n\t\tr%d = 1\n\t} else {\n\t\tr%d = 0\n\t}", i.a, i.b, i.c, i.c)
	case "eqrr":
		return fmt.Sprintf("if r%d == r%d {\n\t\tr%d = 1\n\t} else {\n\t\tr%d = 0\n\t}", i.a, i.b, i.c, i.c)
	case "gtir":
		return fmt.Sprintf("if %d >= r%d {\n\t\tr%d = 1\n\t} else {\n\t\tr%d = 0\n\t}", i.a, i.b, i.c, i.c)
	case "gtrr":
		return fmt.Sprintf("if r%d >= r%d {\n\t\tr%d = 1\n\t} else {\n\t\tr%d = 0\n\t}", i.a, i.b, i.c, i.c)
	case "muli":
		return fmt.Sprintf("r%d = r%d * %d", i.c, i.a, i.b)
	case "seti":
		return fmt.Sprintf("r%d = %d", i.c, i.a)
	case "setr":
		return fmt.Sprintf("r%d = r%d", i.c, i.a)
	default:
		return i.desc
	}
}

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	p := parseInput(f)

	out, err := os.Create("main.go")

	if err != nil {
		panic(err)
	}

	transpiled := p.transpile()

	if _, err := out.Write(transpiled); err != nil {
		panic(err)
	}
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
			instruction{
				desc: fields[0],
				a:    atoi(fields[1]),
				b:    atoi(fields[2]),
				c:    atoi(fields[3]),
			})
	}

	return &res
}

func atoi(a string) int {
	n, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return n
}
