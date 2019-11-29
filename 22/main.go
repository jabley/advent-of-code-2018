package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type input struct {
	depth, x, y int
}

const depth = 11394
const targetX, targetY = 7, 701
const buffer = 150
const printPath = false

const erosionModulus = 20183

const (
	nothing int = iota
	torch
	climbingGear
)

var tileChars = []string{".", "=", "|", "#"}
var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

var weights [targetX + buffer][targetY + buffer][3]int

type entry struct {
	x, y     int
	gear     int
	previous *entry
}

func (e *entry) weight() int {
	return weights[e.x][e.y][e.gear]
}

func parse(r io.Reader) *input {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	scanner.Scan()

	depth := parseInt(scanner.Text())

	scanner.Scan()
	scanner.Scan()

	coords := strings.Split(scanner.Text(), ",")
	x := parseInt(coords[0])
	y := parseInt(coords[1])

	// fmt.Printf("Found depth %v\n", depth)
	// fmt.Printf("Found x %v\n", x)
	// fmt.Printf("Found y %v\n", y)

	return &input{
		depth: depth,
		x:     x,
		y:     y,
	}
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(strings.TrimLeft(s, " "), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func buildMap(constraints *input) [][]int {
	geologicIndex := make([][]int, constraints.x+buffer)
	res := make([][]int, constraints.x+buffer)
	for x := range geologicIndex {
		geologicIndex[x] = make([]int, constraints.y+buffer)
		res[x] = make([]int, constraints.y+buffer)
	}

	for x := 0; x < len(geologicIndex); x++ {
		geologicIndex[x][0] = x * 16807
		res[x][0] = (geologicIndex[x][0] + constraints.depth) % erosionModulus
		for y := 1; y < len(geologicIndex[x]); y++ {
			if x == 0 {
				geologicIndex[0][y] = y * 48271
				res[0][y] = (geologicIndex[0][y] + constraints.depth) % erosionModulus
			} else if x == constraints.x && y == constraints.y {
				res[x][y] = constraints.depth % erosionModulus
			} else {
				geologicIndex[x][y] = res[x-1][y] * res[x][y-1]
				res[x][y] = (geologicIndex[x][y] + constraints.depth) % erosionModulus
			}
		}
	}

	return res
}

func calculateRiskLevel(tiles [][]int, constraints *input) int {
	var res int
	for x := 0; x < len(tiles); x++ {
		for y := 0; y < len(tiles[x]); y++ {
			tiles[x][y] %= 3
			if x <= constraints.x && y <= constraints.y {
				res += tiles[x][y]
			}
		}
	}

	return res
}

func calculateShortestTime(tiles [][]int, target *input) int {
	queue := []*entry{
		&entry{0, 0, torch, nil},
	}

	// You start with the torch equipped.
	weights[0][0][torch] = 1

	for len(queue) > 0 {
		// Find the minimum weight
		min := -1
		for _, entry := range queue {
			if min < 0 || entry.weight() < min {
				min = entry.weight()
			}
		}

		minCount := 0
		for i, currentEntry := range queue {
			if currentEntry.weight() == min {
				if currentEntry.x == target.x && currentEntry.y == target.y && currentEntry.gear == torch {
					return currentEntry.weight() - 1
				}

				for _, dir := range directions {
					x := currentEntry.x + dir[0]
					y := currentEntry.y + dir[1]
					if x < 0 || y < 0 {
						continue
					} else if x >= len(tiles) || y >= len(tiles[x]) {
						continue
					}

					if currentEntry.gear != tiles[x][y] {
						// Current gear can be used on this tile
						posWeight := &weights[x][y][currentEntry.gear]
						if *posWeight == 0 || *posWeight > currentEntry.weight()+1 {
							*posWeight = currentEntry.weight() + 1
							queue = append(queue, &entry{x, y, currentEntry.gear, currentEntry})
						}
					}
				}
				// Try switching gear
				gear := 3 ^ (currentEntry.gear ^ tiles[currentEntry.x][currentEntry.y])
				posWeight := &weights[currentEntry.x][currentEntry.y][gear]
				if *posWeight == 0 || *posWeight > currentEntry.weight()+7 {
					*posWeight = currentEntry.weight() + 7
					queue = append(queue, &entry{currentEntry.x, currentEntry.y, gear, currentEntry})
				}

				// Move the processed entries to the beginning so they can be quickly removed
				queue[minCount], queue[i] = queue[i], queue[minCount]
				minCount++
			}
		}
		queue = queue[minCount:]
	}

	panic("Unable to find result")
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	constraints := parse(f)

	tiles := buildMap(constraints)

	riskLevel := calculateRiskLevel(tiles, constraints)

	fmt.Println("Part A:", riskLevel)

	shortestTime := calculateShortestTime(tiles, constraints)

	fmt.Println("Part B:", shortestTime)
}
