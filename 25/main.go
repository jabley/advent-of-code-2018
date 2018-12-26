package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// point4D is a coordinate in space-time.
type point4D struct {
	x, y, z, t int
}

// manhattanDistance returns the Manhattan distance to the other point.
func (p *point4D) manhattanDistance(other point4D) int {
	return abs(p.x, other.x) + abs(p.y, other.y) + abs(p.z, other.z) + abs(p.t, other.t)
}

type constellation struct {
	points map[point4D]struct{}
	active bool
}

// newConstellation creates a new constellation ready for use.
func newConstellation() *constellation {
	return &constellation{
		points: make(map[point4D]struct{}),
		active: true,
	}
}

// add adds the specified point to this constellation.
func (c *constellation) add(pt point4D) {
	c.points[pt] = struct{}{}
}

// empty removes all points from this constellation.
func (c *constellation) empty() {
	c.points = make(map[point4D]struct{})
}

// merge puts all of the points from the other constellation into this one.
func (c *constellation) merge(other *constellation) {
	for pt := range other.points {
		c.points[pt] = struct{}{}
	}
}

// shouldContain returns true if this constellation contains a point close enough to the given point, otherwise false.
func (c *constellation) shouldContain(pt point4D) bool {
	for p := range c.points {
		if p.manhattanDistance(pt) <= 3 {
			return true
		}
	}

	return false
}

// canJoin returns true if the constellations should contain a point from the other constellation, otherwise false.
func (c *constellation) canJoin(other *constellation) bool {
	for otherPt := range other.points {
		if c.shouldContain(otherPt) {
			return true
		}
	}

	return false
}

// parsePoints parses the non-empty lines in the input into a slice of point4D space-time coordinates.
func parsePoints(r io.Reader) []point4D {

	res := make([]point4D, 0)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		p := point4D{
			// distances: make(map[*point4D]int),
		}

		n, _ := fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.t)

		if n == 4 {
			res = append(res, p)
		}
	}

	return res
}

// countConstellations returns the number of constellations defined by the space-time points.
func countConstellations(points []point4D) int {
	cons := []*constellation{}

nextPoint:
	for _, p := range points {
		for _, c := range cons {
			if c.shouldContain(p) {
				c.add(p)
				// fmt.Printf("Added %v to %v\n", p, c)
				continue nextPoint
			}
		}
		// fmt.Printf("Created new constellation for %v\n", p)
		c := newConstellation()
		c.add(p)
		cons = append(cons, c)
	}

	unstable := true

	for unstable {
		unstable = false
		for i, first := range cons {
			for j, second := range cons {
				if j <= i {
					continue
				}
				if first != second && first.canJoin(second) {
					// fmt.Printf("Merging %v and %v\n", first, second)
					second.merge(first)
					first.active = false
					first.empty()
					unstable = true
				}
			}
		}
	}

	active := []*constellation{}

	for _, c := range cons {
		if c.active {
			active = append(active, c)
		}
	}

	return len(active)
}

// abs returns the absolute difference between 2 ints.
func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Printf("Part 1: %v in %v\n", countConstellations(parsePoints(f)), time.Since(start))
}
