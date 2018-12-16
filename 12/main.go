package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Pots struct {
	state       map[int]string
	_livePlants *[]int
}

func (p *Pots) Min() (i int) {
	for _, k := range p.LivePlants() {
		if k < i {
			i = k
		}
	}
	return
}

func (p *Pots) Max() (i int) {
	for _, k := range p.LivePlants() {
		if k > i {
			i = k
		}
	}
	return
}

func (p *Pots) Sum() (i int) {
	for _, k := range p.LivePlants() {
		i += k
	}
	return
}

func (p *Pots) LivePlants() []int {
	if p._livePlants != nil {
		return *p._livePlants
	}
	res := make([]int, 0)
	for k, v := range p.state {
		if v == "#" {
			res = append(res, k)
		}
	}
	p._livePlants = &res
	return *p._livePlants
}

func (p *Pots) NextGeneration(rules map[string]string) Pots {
	min := p.Min()
	max := p.Max()

	newState := make(map[int]string)

	potState := "...." + p.String() + "...."

	// fmt.Printf("Growing <%v>\n", potState)

	for i := min - 2; i <= max+2; i++ {
		offset := i - (min - 2)
		foo := potState[offset : offset+5]
		if action, ok := rules[foo]; ok && action == "#" {
			// handle the explicit pattern
			// fmt.Printf("Found %v => %v at %d\n", foo, action, i)
			newState[i] = "#"
			continue
		}
		newState[i] = "."
	}
	return Pots{
		state: newState,
	}
}

func (p Pots) String() string {
	var sb strings.Builder
	for i, n := p.Min(), p.Max(); i < n+1; i++ {
		if s, ok := p.state[i]; ok {
			sb.WriteString(s)
		} else {
			sb.WriteString(".")
		}
	}
	return sb.String()
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	pots, rules := ParseInput(f)

	start := time.Now()

	for i := 1; i < 21; i++ {
		pots = pots.NextGeneration(rules)
	}

	fmt.Printf("Part 1 %v in %v\n", pots.Sum(), time.Since(start))
}

func ParseInput(r io.Reader) (pots Pots, rules map[string]string) {
	scanner := bufio.NewScanner(r)

	if scanner.Scan() {
		potState := scanner.Text()[len("initial state: "):]
		pots = parseInitialState(potState)

		scanner.Scan() // skip empty line

		rules = make(map[string]string)

		for scanner.Scan() {
			// parse the rules
			parts := strings.Split(scanner.Text(), " => ")
			rules[parts[0]] = parts[1]
		}
	}

	return
}

func parseInitialState(potState string) Pots {
	tokens := strings.Split(potState, "")

	res := Pots{
		state: make(map[int]string),
	}

	for i, t := range tokens {
		res.state[i] = t
	}

	return res
}
