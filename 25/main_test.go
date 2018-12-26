package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	tests := []struct {
		input           string
		nConstellations int
	}{
		{`
0,0,0,0
3,0,0,0
0,3,0,0
0,0,3,0
0,0,0,3
0,0,0,6
9,0,0,0
12,0,0,0
`,
			2,
		},
		{
			`
-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0
`,

			4,
		},
		{
			`
1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2
`,
			3,
		},
		{
			`
1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2
`,
			8,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%s`, tt.input), func(t *testing.T) {
			points := parsePoints(strings.NewReader(tt.input))
			nConstellations := countConstellations(points)
			if nConstellations != tt.nConstellations {
				t.Errorf("Expected %v but was %v", tt.nConstellations, nConstellations)
			}
		})
	}
}
