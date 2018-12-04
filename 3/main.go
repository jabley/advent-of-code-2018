package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type coord struct{ x, y int }

// Grid is a map of coordinates with a slice containing the claim IDs occupying that coordinate.
type Grid map[coord][]int

type Claim struct {
	id   int
	x, y int
	h, w int
}

func (c *Claim) Right() int {
	return c.x + c.w
}

func (c *Claim) Bottom() int {
	return c.y + c.h
}

func ParseClaims(input string) []Claim {
	regex := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	res := regex.FindAllStringSubmatch(input, -1)

	claims := make([]Claim, 0)

	for _, m := range res {
		claims = append(claims, Claim{
			id: parseInt(m[1]),
			x:  parseInt(m[2]),
			y:  parseInt(m[3]),
			w:  parseInt(m[4]),
			h:  parseInt(m[5]),
		})
	}

	return claims
}

func MakeGrid(claims []Claim) (grid Grid) {
	grid = make(map[coord][]int)
	for _, claim := range claims {
		for x, w := claim.x, claim.Right(); x < w; x++ {
			for y, h := claim.y, claim.Bottom(); y < h; y++ {
				grid[coord{x, y}] = append(grid[coord{x, y}], claim.id)
			}
		}
	}

	return
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return int(i)
}

func CountOverlappingClaimInches(grid Grid) (res int) {
	// How many square inches of fabric are within two or more claims?
	for _, cell := range grid {
		if len(cell) > 1 {
			res++
		}
	}

	return res
}

func FindNonOverlappingClaim(grid Grid, claims []Claim) (res int) {
next_claim:
	for _, claim := range claims {
		for x, w := claim.x, claim.Right(); x < w; x++ {
			for y, h := claim.y, claim.Bottom(); y < h; y++ {
				if len(grid[coord{x, y}]) > 1 {
					continue next_claim
				}
			}
		}

		return claim.id
	}

	return
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	input, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	claims := ParseClaims(string(input))

	grid := MakeGrid(claims)

	// fmt.Println(len(claims))
	fmt.Println(CountOverlappingClaimInches(grid))
	fmt.Println(FindNonOverlappingClaim(grid, claims))
}
