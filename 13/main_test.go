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
		expectedCarts int
		firstCrash    *Coord
	}{
		{`
|
v
|
|
|
^
|
`,
			7, 1,
			2,
			&Coord{x: 0, y: 3},
		},
		{
			`
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   
`,
			6, 13,
			2,
			&Coord{7, 3},
		},
		{
			`
/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
`,
			7, 7,
			9,
			&Coord{2, 0},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%s`, tt.input), func(t *testing.T) {
			tracks := ParseTracks(strings.NewReader(tt.input))
			if len(tracks) != tt.height {
				t.Errorf("Expected height to be %d but was %d", tt.height, len(tracks))
			}
			if len(tracks[0]) != tt.width {
				t.Errorf("Expected width to be %d but was %d", tt.width, len(tracks[0]))
			}

			carts := FindCarts(tracks)

			if len(carts) != tt.expectedCarts {
				t.Errorf("Expected width to be %d but was %d", tt.expectedCarts, len(carts))
			}

			firstCrash := Tick(tracks, carts)

			if !(firstCrash == nil && tt.firstCrash == nil) &&
				((firstCrash == nil && tt.firstCrash != nil) ||
					(firstCrash != nil && tt.firstCrash == nil) ||
					(*firstCrash != *tt.firstCrash)) {
				t.Errorf("Expected firstCrash to be %v but was %v", tt.firstCrash, firstCrash)
			}
		})
	}
}
