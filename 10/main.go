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
)

type pair struct {
	x, y int
}

type point struct {
	pos pair
	d   pair
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	grid, secs := RunSimulation(f)

	fmt.Printf("\n%s\n%d\n", grid, secs)
}

func RunSimulation(r io.Reader) (string, int) {
	points := parsePoints(r)

	minArea := math.MaxInt64

	min, max := pair{}, pair{}

	var t int

	// We could look at the input to see at what time the points are likely to
	// be close, but do the whole thing for now.
	for sec := 1; sec < 11000; sec++ {
		minX, maxX, minY, maxY := math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
		for _, p := range points {
			x := float64(p.pos.x + sec*p.d.x)
			y := float64(p.pos.y + sec*p.d.y)

			// fmt.Printf("x: %f, y: %f\n", x, y)

			minX = math.Min(minX, x)
			maxX = math.Max(maxX, x)
			minY = math.Min(minY, y)
			maxY = math.Max(maxY, y)
		}

		gridSize := gridArea(minX, maxX, minY, maxY)

		if gridSize == 0 {
			fmt.Printf("Got zero grid at time %d\n", sec)
			continue
		}

		if gridSize < minArea {
			t = sec
			minArea = gridSize
			min.x = int(minX)
			min.y = int(minY)
			max.x = int(maxX)
			max.y = int(maxY)
		}
	}

	return fmt.Sprintf("%s", printGrid(points, t, min.x, max.x, min.y, max.y)), t
}

func gridArea(minX, maxX, minY, maxY float64) int {
	x := int(math.Abs(maxX - minX))
	y := int(math.Abs(maxY - minY))

	if x == 0 || y == 0 {
		panic(fmt.Sprintf("got zero length dimension: %f, %f, %f, %f", minX, maxX, minY, maxY))
	}

	return x * y
}

func parsePoints(r io.Reader) []point {
	scanner := bufio.NewScanner(r)
	res := make([]point, 0)

	regex := regexp.MustCompile(`position=<(-? ?\d+), (-? ?\d+)> velocity=<(-? ?\d+), (-? ?\d+)>`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)

		if matches != nil {
			p := point{}
			for _, m := range matches {
				p.pos.x = parseInt(m[1])
				p.pos.y = parseInt(m[2])
				p.d.x = parseInt(m[3])
				p.d.y = parseInt(m[4])
			}

			res = append(res, p)
		}
	}

	return res
}

func printGrid(points []point, t int, minX, maxX, minY, maxY int) string {
	width := int(math.Abs(float64(maxX-minX)) + 1)
	height := int(math.Abs(float64(maxY-minY)) + 1)
	grid := make([][]rune, height)

	for col := range grid {
		grid[col] = make([]rune, width)
		for cell := range grid[col] {
			grid[col][cell] = ' '
		}
	}

	// fmt.Printf("%d, %d, %f, %f\n", len(grid), len(grid[0]), maxX, maxY)

	for _, p := range points {
		x := (p.pos.x + p.d.x*t) - minX
		y := (p.pos.y + p.d.y*t) - minY
		// fmt.Printf("highlighting %d, %d\n", x, y)
		grid[y][x] = '#'
	}

	rows := make([]string, height)

	for i, row := range grid {
		rows[i] = string(row)
	}

	return strings.Join(rows, "\n")
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(strings.TrimLeft(s, " "), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
