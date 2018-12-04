package main

import (
	"testing"
)

func TestSampleProblem(t *testing.T) {

	input := `
#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
`

	claims := ParseClaims(input)

	if len(claims) != 3 {
		t.Errorf("Expected %d claims. Actual: %d", 3, len(claims))
	}

	if claims[0].id != 1 {
		t.Errorf("Expected claim ID %d. Actual: %d", 1, claims[0].id)
	}

	if claims[0].x != 1 {
		t.Errorf("Expected claim x %d. Actual: %d", 1, claims[0].x)
	}

	if claims[0].y != 3 {
		t.Errorf("Expected claim y %d. Actual: %d", 3, claims[0].y)
	}

	if claims[0].w != 4 {
		t.Errorf("Expected claim w %d. Actual: %d", 4, claims[0].w)
	}

	if claims[0].h != 4 {
		t.Errorf("Expected claim h %d. Actual: %d", 4, claims[0].h)
	}

	grid := MakeGrid(claims)

	areaCount := CountOverlappingClaimInches(grid)

	if areaCount != 4 {
		t.Errorf("expected overlapping inches of %d but was %d", 4, areaCount)
	}

	nonOverlapping := FindNonOverlappingClaim(grid, claims)

	if nonOverlapping == 0 {
		t.Errorf("Couldn't find non-overlapping claim")
	}

	if nonOverlapping != 3 {
		t.Errorf("Expected non-overlapping claim ID to be %d but was %d", 3, nonOverlapping)
	}
}
