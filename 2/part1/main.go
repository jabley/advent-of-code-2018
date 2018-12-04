package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	two := 0
	three := 0

	ids := []*ID{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		id := new(line)

		if id.duplicate {
			two++
		}
		if id.triplicate {
			three++
		}
		ids = append(ids, id)
	}

	fmt.Printf("checksum: %d\n", two*three)

	for i, id := range ids {
		matching := hammingDistanceMatch(id, i, ids, 1)
		if matching != nil {
			fmt.Printf("common letters: %s", commonRunes(id.runes, matching.runes))
			break
		}
	}
}

// ID is the box ID
type ID struct {
	runes string
	sums  map[rune]int

	duplicate  bool
	triplicate bool
}

func new(runes string) *ID {
	id := &ID{
		runes: runes,
		sums:  make(map[rune]int),
	}

	for _, r := range runes {
		id.sums[r]++
	}

	for _, v := range id.sums {
		if v == 2 {
			id.duplicate = true
		}
		if v == 3 {
			id.triplicate = true
		}
	}
	return id
}

func hammingDistanceMatch(id *ID, i int, ids []*ID, distance int) *ID {
	if i+1 >= len(ids) {
		return nil
	}

	for _, other := range ids[i+1:] {
		if hammingDistance(id.runes, other.runes) == distance {
			return other
		}
	}

	return nil
}

func hammingDistance(a, b string) int {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Undefined for strings of unequal length. Got <%s> and <%s>", a, b))
	}

	res := 0

	bRunes := []rune(b)

	for i, r := range a {
		if bRunes[i] != r {
			res++
		}
	}

	return res
}

func commonRunes(a, b string) string {
	var res strings.Builder

	bRunes := []rune(b)

	for i, r := range a {
		if bRunes[i] == r {
			res.WriteRune(r)
		}
	}

	return res.String()
}
