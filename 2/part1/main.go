package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	two := 0
	three := 0

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
	}

	fmt.Println(two * three)
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
