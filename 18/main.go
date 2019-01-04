package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type grid [][]rune

func (g grid) String() string {
	rows := []string{}

	for _, r := range g {
		rows = append(rows, string(r))
	}

	return strings.Join(rows, "\n")
}

func (g *grid) neighbours(y, x int) (trees int, lumberyards int) {
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			if !(j >= 0 && j < len(*g)) {
				continue
			}
			if !(i >= 0 && i < len((*g)[0])) {
				continue
			}
			if j == y && i == x {
				// self isn't a neighbour
				continue
			}
			switch (*g)[j][i] {
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}

	return
}

func (g *grid) resourceValue() int {
	trees, lumberyards := g.status()
	return trees * lumberyards
}

func (g *grid) status() (trees int, lumberyards int) {
	for _, row := range *g {
		for _, cell := range row {
			switch cell {
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}
	return
}

func (g grid) tick() grid {
	newGrid := make(grid, len(g))

	for i := range newGrid {
		newGrid[i] = make([]rune, len((g)[0]))
	}

	for y := 0; y < len(g); y++ {
		for x := 0; x < len((g)[0]); x++ {
			trees, lumberyards := g.neighbours(y, x)
			if (g)[y][x] == '.' && trees >= 3 {
				newGrid[y][x] = '|'
			} else if (g)[y][x] == '|' && lumberyards >= 3 {
				newGrid[y][x] = '#'
			} else if (g)[y][x] == '#' && !(lumberyards >= 1 && trees >= 1) {
				newGrid[y][x] = '.'
			} else {
				newGrid[y][x] = (g)[y][x]
			}

			// fmt.Printf("(%d, %d): %c -> %c\n", y, x, g[y][x], newGrid[y][x])
		}
	}

	return newGrid
}

func parseGrid(r io.Reader) grid {
	res := make(grid, 0)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		res = append(res, []rune(scanner.Text()))
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

	g := parseGrid(f)

	// Map of the previously seen patterns, and the time/tick-count when they were seen
	seenStates := make(map[string]int)

	for t := 0; t < 10; t++ {
		g = g.tick()
		seenStates[g.String()] = t
	}

	fmt.Printf("Part 1: %v in %v\n", g.resourceValue(), time.Since(start))

	end := 1000000000

	for t := 10; t < end; t++ {
		g = g.tick()
		if when, seen := seenStates[g.String()]; seen {
			// got a repeating pattern. Skip forward as much as we can.
			skip := (end - t) / (t - when)
			t += skip * (t - when)
		} else {
			seenStates[g.String()] = t
		}
	}

	fmt.Printf("Part 2: %v in %v\n", g.resourceValue(), time.Since(start))
}
