package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

type Direction int

type Turn int

type Coord struct {
	x, y int
}

const (
	_ Direction = iota
	up
	down
	left
	right
)

const occupied = '#'

const (
	leftTurn Turn = iota
	straight
	rightTurn
)

type Cart struct {
	x, y      int
	Direction Direction
	track     rune // the underlying bit of track
	crashed   bool
	turn      Turn
}

func (c *Cart) Move() {
	switch c.Direction {
	case up:
		c.y--
	case down:
		c.y++
	case left:
		c.x--
	case right:
		c.x++
	}
}

func (c *Cart) UpdateDirection(cell rune) {
	switch cell {
	case '/':
		switch c.Direction {
		case left:
			c.Direction = down
		case right:
			c.Direction = up
		case up:
			c.Direction = right
		case down:
			c.Direction = left
		}
	case '\\':
		switch c.Direction {
		case left:
			c.Direction = up
		case right:
			c.Direction = down
		case up:
			c.Direction = left
		case down:
			c.Direction = right
		}
	case '+':
		c.Turn()
	}
}

func (c *Cart) Turn() {
	switch c.Direction {
	case up:
		switch c.turn {
		case leftTurn:
			c.Direction = left
			c.turn = straight
		case straight:
			c.Direction = up
			c.turn = rightTurn
		case rightTurn:
			c.Direction = right
			c.turn = leftTurn
		}
	case down:
		switch c.turn {
		case leftTurn:
			c.Direction = right
			c.turn = straight
		case straight:
			c.Direction = down
			c.turn = rightTurn
		case rightTurn:
			c.Direction = left
			c.turn = leftTurn
		}
	case left:
		switch c.turn {
		case leftTurn:
			c.Direction = down
			c.turn = straight
		case straight:
			c.Direction = left
			c.turn = rightTurn
		case rightTurn:
			c.Direction = up
			c.turn = leftTurn
		}
	case right:
		switch c.turn {
		case leftTurn:
			c.Direction = up
			c.turn = straight
		case straight:
			c.Direction = right
			c.turn = rightTurn
		case rightTurn:
			c.Direction = down
			c.turn = leftTurn
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tracks := ParseTracks(f)
	carts := FindCarts(tracks)

	start := time.Now()

	firstCrash := Tick(tracks, carts)

	fmt.Printf("Took %v\n", time.Since(start))
	fmt.Printf("part 1: %d,%d\n", firstCrash.x, firstCrash.y)
}

func ParseTracks(r io.Reader) [][]rune {
	scanner := bufio.NewScanner(r)

	res := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()

		// Skip blank lines
		if len(line) == 0 {
			continue
		}

		res = append(res, []rune(line))
	}

	return res
}

func FindCarts(tracks [][]rune) (carts []Cart) {
	for y := 0; y < len(tracks); y++ {
		for x := 0; x < len(tracks[0]); x++ {
			var d Direction
			var cell rune
			switch tracks[y][x] {
			case '^':
				d = up
				cell = '|'
			case 'v':
				d = down
				cell = '|'
			case '<':
				d = left
				cell = '-'
			case '>':
				d = right
				cell = '-'
			}
			if cell != 0 {
				carts = append(carts, Cart{x: x, y: y, Direction: d, track: cell})
				tracks[y][x] = occupied
			}
		}
	}

	return
}

func Tick(tracks [][]rune, carts []Cart) (firstCrash *Coord) {
	nCarts := len(carts)

	for t := 1; nCarts > 1; t++ {
		// Carts on the first row move first (acting from left to right) ...
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})

	nextCart:
		for i := range carts {
			c := &carts[i]

			// fmt.Printf("Moving cart at %d,%d\n", c.x, c.y)

			if !c.crashed {

				// Mark the track as unoccupied
				tracks[c.y][c.x] = c.track

				// Move
				c.Move()

				// Collision detection?
				if tracks[c.y][c.x] == occupied {
					// fmt.Printf("crash: %d,%d\n", c.x, c.y)
					if firstCrash == nil {
						firstCrash = &Coord{x: c.x, y: c.y}
					}
					c.crashed = true

					// Find the other cart in this location. This is the cart which set the occupied rune
					for i, cc := range carts {
						if cc.x == c.x && cc.y == c.y && !cc.crashed {
							carts[i].crashed = true
							// Restore the track
							tracks[cc.y][cc.x] = cc.track
							break
						}
					}

					nCarts -= 2
					continue nextCart
				}

				// Update cart state based on track layout
				c.UpdateDirection(tracks[c.y][c.x])

				// push the track state into the cart and mark this cell of the track as occupied
				c.track = tracks[c.y][c.x]
				tracks[c.y][c.x] = occupied
			}
		}

		// fmt.Printf("nr of carts after tick %d: %d\n", t, nCarts)
	}

	return
}
