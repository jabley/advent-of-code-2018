package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type bot struct {
	x, y, z, r int
}

func (b bot) withinRange(controller bot) bool {
	return manhattanDistance(b, controller) <= controller.r
}

func manhattanDistance(a, b bot) int {
	return abs(a.x, b.x) + abs(a.y, b.y) + abs(a.z, b.z)
}

// abs returns the absolute difference between 2 ints.
func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	bots := parse(f)
	strongest := findStrongest(bots)
	partA := withinRange(bots, strongest)

	fmt.Println("Part A:", len(partA))
}

func parse(r io.Reader) []bot {
	res := make([]bot, 0)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		b := bot{}

		n, _ := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &b.x, &b.y, &b.z, &b.r)

		if n == 4 {
			res = append(res, b)
		}
	}

	return res
}

func findStrongest(bots []bot) bot {
	max := 0

	var res bot

	for _, b := range bots {
		if b.r > max {
			max = b.r
			res = b
		}
	}

	return res
}

func withinRange(bots []bot, strongest bot) []bot {
	var res []bot

	for _, b := range bots {
		if b.withinRange(strongest) {
			res = append(res, b)
		}
	}

	return res
}
