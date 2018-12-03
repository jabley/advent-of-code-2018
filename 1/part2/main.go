package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	lines := readLines(f)

	var currentFrequency int64

	seenFrequencies := make(map[int64]bool)

	var duplicate int64

	seenFrequencies[currentFrequency] = true

quit:
	for {
		for _, line := range lines {
			change, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}
			currentFrequency += change

			if _, contains := seenFrequencies[currentFrequency]; contains {
				duplicate = currentFrequency
				break quit
			} else {
				seenFrequencies[currentFrequency] = true
			}
		}
	}

	fmt.Println(duplicate)
}

func readLines(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
