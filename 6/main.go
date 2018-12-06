package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type coord struct {
	x, y int
}

var offGrid = coord{-1, -1}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	max, regionSize := CalculateArea(f, 10000)
	fmt.Printf("%d\n", max)
	fmt.Printf("%d\n", regionSize)
}

func CalculateArea(r io.Reader, maxDistance int) (int, int) {
	points, size := parseCoords(r)

	// Set of infinite coordinates
	infiniteCoords := make(map[coord]struct{})

	// Map of number of locations that each coordinate owns
	coordinateAreas := make(map[coord]int)

	regionSize := 0

	for y := 0; y <= size.y; y++ {
		for x := 0; x <= size.x; x++ {
			minDistance := math.MaxInt64
			closestCoord := offGrid
			totalDistance := 0

			for _, c := range points {
				distance := calculateDistance(c, x, y)
				totalDistance += distance

				if distance < minDistance {
					minDistance = distance
					closestCoord = c
				} else if distance == minDistance {
					closestCoord = offGrid
				}
			}

			if x == 0 || x == size.x || y == 0 || y == size.y {
				infiniteCoords[closestCoord] = struct{}{}
			}

			if closestCoord != offGrid {
				coordinateAreas[closestCoord]++
			}

			if totalDistance < maxDistance {
				regionSize++
			}
		}
	}

	max := 0
	for c, n := range coordinateAreas {
		if _, isInfinite := infiniteCoords[c]; n > max && !isInfinite {
			max = n
		}
	}

	return max, regionSize
}

func calculateDistance(c coord, x int, y int) (res int) {
	if c.x < x {
		res += x - c.x
	} else {
		res += c.x - x
	}

	if c.y < y {
		res += y - c.y
	} else {
		res += c.y - y
	}

	return
}

// Returns the points, plus the area of the grid
func parseCoords(r io.Reader) ([]coord, coord) {
	res := make([]coord, 0)
	size := coord{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		point := coord{}
		fmt.Sscanf(line, "%d, %d", &point.x, &point.y)
		res = append(res, point)
		if point.x > size.x {
			size.x = point.x
		}
		if point.y > size.y {
			size.y = point.y
		}
	}

	return res, size
}
