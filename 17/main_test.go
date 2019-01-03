package main

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	input := `x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504`

	//   44444455555555
	//   99999900000000
	//   45678901234567
	// 0 ......+.......
	// 1 ............#.
	// 2 .#..#.......#.
	// 3 .#..#..#......
	// 4 .#..#..#......
	// 5 .#.....#......
	// 6 .#.....#......
	// 7 .#######......
	// 8 ..............
	// 9 ..............
	// 10 ....#.....#...
	// 11 ....#.....#...
	// 12 ....#.....#...
	// 13 ....#######...

	clay := parseScan(strings.NewReader(input))
	if len(clay) != 34 {
		t.Fatalf("Umable to parse clay. Expected %d but got %d", 34, len(clay))
	}

	state := newState(clay)

	if state.maxY != 13 {
		t.Errorf("expected maxY of %d but was %d", 13, state.maxY)
	}

	state.fill(point{x: 500, y: 0}, dirDown)

	wet, retained := state.report()

	if wet != 57 {
		t.Fatalf("Expected %d wet tiles but was %d", 57, wet)
	}

	if retained != 29 {
		t.Fatalf("Expected %d tiles to retain water, but was %d", 29, retained)
	}
}
