package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

type kind int

const (
	kindWall kind = iota
	kindSpace
	kindGoblin
	kindElf
)

var runeKinds = map[rune]kind{
	'.': kindSpace,
	'E': kindElf,
	'G': kindGoblin,
	'#': kindWall,
}

var squareRunes = map[kind]rune{
	kindSpace:  '.',
	kindElf:    'E',
	kindGoblin: 'G',
	kindWall:   '#',
}

var unitKinds = map[kind]string{
	kindElf:    "Elf",
	kindGoblin: "Goblin",
}

const (
	defaultAP = 3
	defaultHP = 200
)

type coord struct {
	x, y int
}

type square struct {
	kind kind
	x, y int
	unit *unit
}

func (s square) String() string {
	return fmt.Sprintf("%d,%d", s.x, s.y)
}

func (s *square) findWalkableSquares(game *game) (map[*square]int, map[*square]*square) {
	// frontier is the squares to be explored
	frontier := []*square{s}
	// distance is a map of squares and the cost of getting to them (number of moves)
	distance := map[*square]int{s: 0}
	// cameFrom is a map of squares and the parent square that led to them
	cameFrom := map[*square]*square{s: nil}

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		for _, neighbour := range current.walkableNeighbours(game) {
			if _, ok := distance[neighbour]; !ok {
				// Not seen this before
				frontier = append(frontier, neighbour)
				distance[neighbour] = distance[current] + 1
				cameFrom[neighbour] = current
			}
		}
	}

	return distance, cameFrom
}

func (s *square) walkableNeighbours(game *game) []*square {
	var neighbours []*square

	for _, point := range compassPoints {
		if n := game.board[s.y+point.y][s.x+point.x]; n != nil && n.kind == kindSpace {
			neighbours = append(neighbours, n)
		}
	}

	// fmt.Printf("%v has walkable neighbours %v\n", s, neighbours)

	return neighbours
}

type readingOrder []*square

func (ro readingOrder) Len() int      { return len(ro) }
func (ro readingOrder) Swap(i, j int) { ro[i], ro[j] = ro[j], ro[i] }
func (ro readingOrder) Less(i, j int) bool {
	if ro[i].y == ro[j].y {
		return ro[i].x < ro[j].x
	}
	return ro[i].y < ro[j].y
}

type unit struct {
	hp     int
	ap     int
	square *square
	kind   kind
}

func (u unit) String() string {
	return fmt.Sprintf("%v (hp: %d, ap: %d) at %v", unitKinds[u.kind], u.hp, u.ap, u.square)
}

type sortableUnits []*unit

func (su sortableUnits) Len() int      { return len(su) }
func (su sortableUnits) Swap(i, j int) { su[i], su[j] = su[j], su[i] }
func (su sortableUnits) Less(i, j int) bool {
	if su[i].square.y == su[j].square.y {
		return su[i].square.x < su[j].square.x
	}
	return su[i].square.y < su[j].square.y
}

type heuristicEnemy struct {
	distance int
	unit     *unit
}

// To attack, the unit first determines all of the targets that are in
// range of it by being immediately adjacent to it. If there are no
// such targets, the unit ends its turn. Otherwise, the adjacent target
//  with the fewest hit points is selected; in a tie, the adjacent
// target with the fewest hit points which is first in reading order is
// selected.
//
// So the order here is important. We should do it in reading order.
var compassPoints = []coord{
	{0, -1}, // N
	{-1, 0}, // W
	{1, 0},  // E
	{0, 1},  // S
}

func (u *unit) attack(enemy *unit) {
	enemy.hp -= u.ap
	if enemy.hp <= 0 {
		// dead
		enemy.square.kind = kindSpace
		enemy.square.unit = nil
		enemy.square = nil
	}
}

func (u *unit) findEnemies(game *game) (res []*unit) {
	for _, a := range game.units {
		if a.kind != u.kind && a.hp > 0 {
			res = append(res, a)
		}
	}
	return
}

// hasAttackOption returns the attackable enemy if N, S, E, or W is occupied by an enemy
//
// To attack, the unit first determines all of the targets that are in
// range of it by being immediately adjacent to it. If there are no
// such targets, the unit ends its turn. Otherwise, the adjacent target
//  with the fewest hit points is selected; in a tie, the adjacent
// target with the fewest hit points which is first in reading order is
// selected.
func (u *unit) hasAttackOption(game *game) *unit {
	var target *unit

	for _, p := range compassPoints {
		if s := game.board[u.square.y+p.y][u.square.x+p.x]; s != nil {
			// There is an adjacent square.
			// Does it contain an attackable enemy?
			if s.unit != nil && s.unit.kind != u.kind && s.unit.hp > 0 {
				if target == nil || target.hp > s.unit.hp {
					// It's the first target in reading order, or it's
					// got lower hp than the previously selected target
					target = s.unit
				}
			}
		}
	}

	return target
}

func (u *unit) move(game *game) {
	step, _ := u.nextStep(game)
	// fmt.Printf("Updating %v to move to %v\n", u, step)
	if step != nil {
		step.kind = u.kind
		step.unit = u
		u.square.kind = kindSpace
		u.square.unit = nil
		u.square = step
	}
}

func (u *unit) nextStep(game *game) (*square, *square) {
	// fmt.Printf("Moving %v\n", u)

	closestTargetDistance := math.MaxInt32

	distances, cameFrom := u.square.findWalkableSquares(game)

	enemies := u.findEnemies(game)

	var targets readingOrder

	// fmt.Printf("%v has enemies %v\n", u, enemies)

	for _, enemy := range enemies {
		for _, target := range enemy.square.walkableNeighbours(game) {
			if distance, ok := distances[target]; ok && distance <= closestTargetDistance {
				if distance < closestTargetDistance {
					closestTargetDistance = distance
					targets = readingOrder{}
				}
				targets = append(targets, target)
			}
		}
	}

	sort.Sort(targets)

	if len(targets) > 0 {
		target := targets[0]
		current := target

		// Walk back through the path from the goal to the start to
		// find the next square
		for {
			if cameFrom[current] == u.square {
				return current, target
			}
			current = cameFrom[current]
		}
	}

	return nil, nil
}

func (u *unit) takeTurn(g *game) {
	if enemy := u.hasAttackOption(g); enemy != nil {
		// fmt.Printf("Unit %v is attacking %v\n", u, enemy)
		u.attack(enemy)
		return
	}

	u.move(g)

	if enemy := u.hasAttackOption(g); enemy != nil {
		u.attack(enemy)
	}
}

type game struct {
	board map[int]map[int]*square
	units []*unit
}

func (g *game) boardHeight() int {
	return len(g.board)
}

func (g *game) boardWidth() int {
	return len(g.board[0])
}

func (g *game) drawCave() string {
	var rows []string

	for i, n := 0, len(g.board); i < n; i++ {
		var sb strings.Builder
		for j, m := 0, len(g.board[i]); j < m; j++ {
			sb.WriteRune(squareRunes[g.board[i][j].kind])
		}
		rows = append(rows, sb.String())
	}

	return strings.Join(rows, "\n")
}
func (g *game) isComplete() bool {
	seenTypes := make(map[kind]struct{})

	for _, u := range g.units {
		if _, ok := seenTypes[u.kind]; !ok {
			seenTypes[u.kind] = struct{}{}
		}
	}

	return len(seenTypes) == 1
}

func (g *game) play() int {
	for i := 1; true; i++ {
		fullRound := g.playRound()

		if !fullRound {
			return g.remainingHp() * (i - 1)
		} else if g.isComplete() {
			return g.remainingHp() * i
		}
	}

	return -1
}

// playRound returns false if the round did not fully complete, otherwise true.
func (g *game) playRound() bool {
	g.removeDead()

	// Sort into reading order of their starting positions in the round
	sort.Sort(sortableUnits(g.units))

	for _, u := range g.units {
		// fmt.Printf("Unit at %d,%d is taking turn\n", u.square.x, u.square.y)
		if u.hp <= 0 {
			// dead, skip it
			continue
		}

		if len(u.findEnemies(g)) == 0 {
			return false
		}

		u.takeTurn(g)
	}

	return true
}

func (g *game) remainingHp() int {
	res := 0

	for _, u := range g.units {
		if u.hp > 0 {
			res += u.hp
		}
	}

	return res
}
func (g *game) removeDead() {
	var newUnits []*unit

	for _, unit := range g.units {
		if unit.hp > 0 {
			newUnits = append(newUnits, unit)
		}
	}

	g.units = newUnits
}

func parseBoard(r io.Reader) *game {
	scanner := bufio.NewScanner(r)

	board := make(map[int]map[int]*square)
	units := make([]*unit, 0)

	for y := 0; scanner.Scan(); y++ {
		row := scanner.Text()

		// Skip blank lines
		if len(row) == 0 {
			y--
			continue
		}

		if _, ok := board[y]; !ok {
			board[y] = make(map[int]*square)
		}

		for x, cell := range row {
			kind, ok := runeKinds[cell]
			if !ok {
				kind = kindWall
			}

			square := &square{kind: kind, x: x, y: y}
			board[y][x] = square

			switch kind {
			case kindGoblin:
				units = append(units, newUnit(kindGoblin, square))
			case kindElf:
				units = append(units, newUnit(kindElf, square))
			}
		}
	}

	return &game{board: board, units: units}
}

func newUnit(kind kind, square *square) *unit {
	res := &unit{
		ap:     defaultAP,
		hp:     defaultHP,
		kind:   kind,
		square: square,
	}

	square.unit = res

	return res
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	start := time.Now()

	game := parseBoard(f)
	outcome := game.play()

	fmt.Printf("%v in %v\n", outcome, time.Since(start))
}
