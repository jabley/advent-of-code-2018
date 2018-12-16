package main

import (
	"fmt"
	"time"
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func main() {
	serial := 6878
	start := time.Now()
	part1 := BiggestClusterFor3(serial)
	duration := time.Since(start)
	fmt.Printf("part 1: %v in %v\n", part1, duration)

	start = time.Now()
	part2 := BiggestAnySizeCluster(serial)
	duration = time.Since(start)
	fmt.Printf("part 2: %v in %v\n", part2, duration)
}

func populateCells(serial int) [][]int {
	cells := make([][]int, 300)
	for y := range cells {
		cells[y] = make([]int, 300)
		for x := range cells[y] {
			rackID := x + 1 + 10
			power := rackID * (y + 1)
			power += serial
			power *= rackID
			power = (power / 100) % 10
			cells[y][x] = power - 5
		}
	}

	return cells
}

func powerfulCluster(cells [][]int, size int) (int, coord) {
	n := len(cells) - size + 1
	clusters := make([][]int, n)
	maxValue := 0
	var res coord

	for y := range clusters {
		clusters[y] = make([]int, n)
		for x := range clusters[0] {
			sum := 0
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					sum += cells[y+j][x+i]
				}
			}
			if sum > maxValue {
				res = coord{x: x + 1, y: y + 1}
				maxValue = sum
			}
		}
	}

	return maxValue, res
}

func BiggestClusterFor3(serial int) coord {
	cells := populateCells(serial)

	_, res := powerfulCluster(cells, 3)

	return res
}

type result struct {
	power int
	size  int
	coord coord
}

func BiggestAnySizeCluster(serial int) string {
	cells := populateCells(serial)

	sumTable := BuildSumTable(cells)

	best := 0
	var c coord
	size := 0

	for s := 1; s < 301; s++ {
		for x := 0; x < 301-s; x++ {
			for y := 0; y < 301-s; y++ {
				power := CalculateSquarePower(sumTable, x, y, s)
				if power > best {
					best = power
					c = coord{x: x + 1, y: y + 1}
					size = s
				}
			}
		}
	}

	return fmt.Sprintf("%s,%d", c, size)
}

func CalculateSquarePower(sumTable [][]int, x, y, s int) int {
	top := y - 1
	left := x - 1
	bottom := y + s - 1
	right := x + s - 1
	var tl, tr, bl, br int

	br = sumTable[bottom][right]

	if top < 0 && left < 0 {
		tr = 0
		tl = 0
		bl = 0
	} else if top < 0 {
		tr = 0
		tl = 0
		bl = sumTable[bottom][left]
	} else if left < 0 {
		tl = 0
		bl = 0
		tr = sumTable[top][right]
	} else {
		tl = sumTable[top][left]
		tr = sumTable[top][right]
		bl = sumTable[bottom][left]
	}

	return br + tl - bl - tr
}

func BuildSumTable(cells [][]int) [][]int {
	res := make([][]int, len(cells))

	for col, y := range cells {
		res[col] = make([]int, len(cells[col]))
		for row, x := range y {
			v := x
			if row != 0 && col != 0 {
				v += res[col-1][row] + res[col][row-1] - res[col-1][row-1]
			} else if row != 0 {
				v += res[col][row-1]
			} else if col != 0 {
				v += res[col-1][row]
			}
			res[col][row] = v
		}
	}

	return res
}
