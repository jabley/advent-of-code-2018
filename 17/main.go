package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

type squareType map[point]struct{}

type phase int

const (
	sand phase = iota
	clay
	flowing
	settled
)

var (
	dirDown  = point{x: 0, y: 1}
	dirLeft  = point{x: -1, y: 0}
	dirRight = point{x: 1, y: 0}
)

type state struct {
	clay    squareType
	settled squareType
	flowing squareType
	minY    int
	maxY    int
}

func (s *state) is(p phase, pt point) (res bool) {
	switch p {
	case clay:
		_, res = s.clay[pt]
	case settled:
		_, res = s.settled[pt]
	case flowing:
		_, res = s.flowing[pt]
	}

	return
}

func (s *state) fill(source point, direction point) bool {
	// fmt.Printf("Filling %v\n", source)
	s.flowing[source] = struct{}{}

	below := point{x: source.x, y: source.y + 1}

	if !s.is(clay, below) && !s.is(flowing, below) && s.withinRange(below.y, 1) {
		s.fill(below, dirDown)
	}

	if !s.is(clay, below) && !s.is(settled, below) {
		// fmt.Printf("Done with %v\n", source)
		return false
	}

	left := point{x: source.x - 1, y: source.y}
	right := point{x: source.x + 1, y: source.y}

	// fmt.Printf("Filling left %v\n", left)
	leftFilled := s.is(clay, left) || !s.is(flowing, left) && s.fill(left, dirLeft)

	// fmt.Printf("Filling right %v\n", left)
	rightFilled := s.is(clay, right) || !s.is(flowing, right) && s.fill(right, dirRight)

	if direction == dirDown && leftFilled && rightFilled {
		s.settled[source] = struct{}{}

		for s.is(flowing, left) {
			s.settled[left] = struct{}{}
			// left.x--
			left = point{x: left.x - 1, y: left.y}
		}

		for s.is(flowing, right) {
			s.settled[right] = struct{}{}
			// right.x++
			right = point{x: right.x + 1, y: right.y}
		}
	}

	return (direction == dirLeft && leftFilled || s.is(clay, left)) ||
		(direction == dirRight && rightFilled || s.is(clay, right))
}

func (s *state) print() {
	minX := minDimension(s.clay, func(pt point) int {
		return pt.x
	})
	maxX := maxDimension(s.clay, func(pt point) int {
		return pt.x
	})

	grid := []string{}
	for y := s.minY - 1; y <= s.maxY; y++ {
		var sb strings.Builder
		for x := minX - 1; x <= maxX+1; x++ {
			tracer := point{x: x, y: y}
			if x == 500 && y == s.minY-1 {
				sb.WriteRune('+')
			} else if s.is(settled, tracer) {
				sb.WriteRune('~')
			} else if s.is(flowing, tracer) {
				sb.WriteRune('|')
			} else if s.is(clay, tracer) {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		grid = append(grid, sb.String())
	}

	fmt.Printf("%s\n", strings.Join(grid, "\n"))
}

func (s *state) report() int {
	s.print()

	wet := make(squareType)

	for cell := range s.flowing {
		if s.withinRange(cell.y, s.minY) {
			wet[cell] = struct{}{}
			// } else {
			// 	fmt.Printf("Ignoring %v\n", cell)
		}
	}

	for cell := range s.settled {
		if s.withinRange(cell.y, s.minY) {
			wet[cell] = struct{}{}
			// } else {
			// 	fmt.Printf("Ignoring %v\n", cell)
		}
	}

	return len(wet)
}

func (s *state) withinRange(y, minY int) bool {
	return y >= minY && y <= s.maxY
}

func newState(clay squareType) *state {
	return &state{
		clay:    clay,
		settled: make(squareType),
		flowing: make(squareType),
		minY:    minDimension(clay, func(pt point) int { return pt.y }),
		maxY:    maxDimension(clay, func(pt point) int { return pt.y }),
	}
}

func maxDimension(clay squareType, dim func(point) int) int {
	max := math.MinInt64

	for square := range clay {
		i := dim(square)
		if max < i {
			max = i
		}
	}

	return max
}

func minDimension(clay squareType, dim func(point) int) int {
	min := math.MaxInt64

	for square := range clay {
		i := dim(square)
		if i < min {
			min = i
		}
	}

	return min
}

func parseScan(r io.Reader) (clay squareType) {
	scanner := bufio.NewScanner(r)

	matcher := regexp.MustCompile(`(x|y)=(\d+), (x|y)=(\d+)..(\d+)`)

	clay = make(squareType)

	for scanner.Scan() {
		matches := matcher.FindAllStringSubmatch(scanner.Text(), -1)
		if matches != nil {
			if matches[0][1] == "x" {
				for y := atoi(matches[0][4]); y <= atoi(matches[0][5]); y++ {
					clay[point{x: atoi(matches[0][2]), y: y}] = struct{}{}
				}
			} else {
				for x := atoi(matches[0][4]); x <= atoi(matches[0][5]); x++ {
					clay[point{x: x, y: atoi(matches[0][2])}] = struct{}{}
				}
			}
		}
	}

	return
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

	state := newState(parseScan(f))

	state.fill(point{x: 500, y: 0}, dirDown)

	wet := state.report()

	fmt.Printf("Took %v\n", time.Since(start))
	fmt.Printf("Part 1: %d\n", wet)
}
