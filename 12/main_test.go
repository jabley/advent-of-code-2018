package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {

	//                       0         1         2
	//                       01234567890123456789012345
	input := `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`

	initialState := "#..#.#..##......###...###"

	expectedRules := map[string]string{
		"...##": "#",
		"..#..": "#",
		".#...": "#",
		".#.#.": "#",
		".#.##": "#",
		".##..": "#",
		".####": "#",
		"#.#.#": "#",
		"#.###": "#",
		"##.#.": "#",
		"##.##": "#",
		"###..": "#",
		"###.#": "#",
		"####.": "#",
	}

	pots, rules := ParseInput(strings.NewReader(input))

	initial := pots

	if pots.String() != initialState {
		t.Errorf("Expected %v but got %v", initialState, pots.String())
	}
	if !reflect.DeepEqual(expectedRules, rules) {
		t.Errorf("Expected %v but got %v", expectedRules, rules)
	}

	tests := []struct {
		livePlants int
		min        int
		max        int
		sum        int
		nextGen    string
	}{
		{
			11, 0, 24, 145,
			`#..#.#..##......###...###`,
		},
		{
			7, 0, 24, 91,
			`#...#....#.....#..#..#..#`,
		},
		{
			11, 0, 25, 132,
			`##..##...##....#..#..#..##`,
		},
		{
			9, -1, 25, 102,
			`#.#...#..#.#....#..#..#...#`,
		},
		{
			11, 0, 26, 154,
			`#.#..#...#.#...#..#..##..##`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf(`%d`, i), func(t *testing.T) {
			if len(pots.LivePlants()) != tt.livePlants {
				t.Errorf("expected live plants %v but was %v", tt.livePlants, pots.LivePlants())
			}
			if pots.Min() != tt.min {
				t.Errorf("expected min %v but was %v", tt.min, pots.Min())
			}

			if pots.Max() != tt.max {
				t.Errorf("expected max %v but was %v", tt.max, pots.Max())
			}
			if pots.Sum() != tt.sum {
				t.Errorf("expected sum %v but was %v", tt.sum, pots.Sum())
			}

			if !strings.Contains(pots.String(), tt.nextGen) {
				t.Errorf("Expected to contain <%v> but got <%v>", tt.nextGen, pots.String())
			}

			pots = pots.NextGeneration(rules)
		})
	}

	for i := 1; i < 21; i++ {
		initial = initial.NextGeneration(rules)
	}

	if initial.Sum() != 325 {
		t.Errorf("Expected 20th generation sum to be 325 but was %d", initial.Sum())
	}
}
