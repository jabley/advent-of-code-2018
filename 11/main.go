package main

import (
	"fmt"
	"sync"
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func main() {
	serial := 6878
	fmt.Printf("%v\n", BiggestClusterFor3(serial))
	fmt.Printf("%v\n", BiggestAnySizeCluster(serial))
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

type result struct {
	power int
	size  int
	coord coord
}

func BiggestAnySizeCluster(serial int) string {
	cells := populateCells(serial)

	best := 0
	var coord coord
	size := 0

	results := make(chan result)
	var wg sync.WaitGroup

	for i := 1; i < 301; i++ {
		wg.Add(1)
		go func(size int) {
			defer wg.Done()
			power, c := powerfulCluster(cells, size)
			results <- result{
				power: power,
				size:  size,
				coord: c,
			}
		}(i)
	}

	go func() {
		for {
			select {
			case res, more := <-results:
				if !more {
					// Channel has been closed, exit this goroutine
					return
				}
				if res.power > best {
					best = res.power
					coord = res.coord
					size = res.size
				}
			}
		}
	}()

	wg.Wait()

	// All results have been consumed, there will be no more results.
	close(results)

	return fmt.Sprintf("%s,%d", coord, size)
}
