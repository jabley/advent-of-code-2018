package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		input         string
		height, width int
		nUnits        int
		// firstUnit     unit
		rounds        int
		expectedBoard string
	}{
		// In range:     Nearest:      Chosen:       Distance:     Step:
		// #######       #######       #######       #######       #######
		// #.E...#       #.E...#       #.E...#       #4E212#       #..E..#
		// #...?.#  -->  #...!.#  -->  #...+.#  -->  #32101#  -->  #.....#
		// #..?G?#       #..!G.#       #...G.#       #432G2#       #...G.#
		// #######       #######       #######       #######       #######
		{`
#######
#.E...#
#.....#
#...G.#
#######`,
			5, 7, 2,
			1,
			`#######
#..E..#
#...G.#
#.....#
#######`,
		},
		{
			`
########
#E...G.#
########
`,
			3,
			8,
			2,
			1,
			`########
#.E.G..#
########`,
		},
		{
			`#########
#G..G..G#
#.......#
#.......#
#G..E..G#
#.......#
#.......#
#G..G..G#
#########`,
			9, 9, 9, 1,
			`#########
#.G...G.#
#...G...#
#...E..G#
#.G.....#
#.......#
#G..G..G#
#.......#
#########`,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%v`, tt.input), func(t *testing.T) {
			g := parseBoard(strings.NewReader(tt.input))
			if g.boardHeight() != tt.height {
				t.Errorf("Expected height of %d but was %d", tt.height, g.boardHeight())
			}

			if g.boardWidth() != tt.width {
				t.Errorf("Expected width of %d but was %d", tt.width, g.boardWidth())
			}

			if len(g.units) != tt.nUnits {
				t.Errorf("Expected %d units but was %d", tt.nUnits, len(g.units))
			}

			// if tt.firstUnit.square.x != g.units[0].square.x {
			// 	t.Errorf("Expected %v but was %v", tt.firstUnit, g.units[0])
			// }

			for i := 0; i < tt.rounds; i++ {
				g.playRound(false)
			}

			cave := g.drawCave()

			if cave != tt.expectedBoard {
				t.Errorf("Expected \n%v\n\nbut was \n%v", tt.expectedBoard, cave)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input   string
		outcome int
	}{
		{
			`
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
`,
			27730,
		},
		{
			`#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`,
			36334,
		},
		{
			`#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			39514,
		},
		{
			`
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			27755,
		},
		{
			`
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######
`,
			28944,
		},
		{
			`
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########			
`,
			18740,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%s`, tt.input), func(t *testing.T) {
			g := parseBoard(strings.NewReader(tt.input))
			outcome := g.play()

			if outcome != tt.outcome {
				t.Errorf("Expected %d but was %d", tt.outcome, outcome)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input   string
		outcome int
	}{
		{
			`
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
`,
			4988,
		},
		{
			`
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			31284,
		},
		{
			`
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			3478,
		},
		{
			`
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######
`,
			6474,
		},
		{
			`
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########
`,
			1140,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%s`, tt.input), func(t *testing.T) {
			outcome := playElfBoost(tt.input)

			if outcome != tt.outcome {
				t.Errorf("Expected %d but was %d", tt.outcome, outcome)
			}
		})
	}
}
