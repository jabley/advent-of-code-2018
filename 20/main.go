package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type point struct {
	x, y int
}

type context struct {
	dist, x, y int
}

func path(r io.Reader) map[point]int {
	grid := make(map[point]int)
	dist, x, y := 0, 0, 0

	stack := []context{}

	buf := bufio.NewReader(r)

	for {
		r, _, err := buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		switch r {
		case '^', '$':
		case '(':
			// push
			stack = append(stack, context{dist: dist, x: x, y: y})
		case ')':
			// pop
			c := stack[len(stack)-1]
			dist, x, y = c.dist, c.x, c.y
			stack = stack[:len(stack)-1]
		case '|':
			// peek
			c := stack[len(stack)-1]
			dist, x, y = c.dist, c.x, c.y
		default:
			switch r {
			case 'E':
				x++
			case 'W':
				x--
			case 'S':
				y++
			case 'N':
				y--
			}
			dist++

			d, ok := grid[point{x: x, y: y}]
			if !ok || dist < d {
				grid[point{x: x, y: y}] = dist
			}
		}
	}

	return grid
}

func max(grid map[point]int) int {
	res := 0
	for _, v := range grid {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	grid := path(f)

	fmt.Printf("Part 1 in %v: %d\n", time.Since(start), max(grid))
}
