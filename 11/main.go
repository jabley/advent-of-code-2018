package main

import "fmt"

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func main() {
	serial := 6878
	fmt.Printf("%v\n", BiggestClusterFor3(serial))
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
		for x := range clusters[y] {
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
