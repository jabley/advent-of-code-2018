package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const charDiff = 'a' - 'A'

func React(src string) int {
	res := []rune{}

	for _, r := range src {
		if len(res) > 0 {
			diff := res[len(res)-1] - r

			if diff == charDiff || diff == -charDiff {
				res = res[0 : len(res)-1]
				continue
			}
		}
		res = append(res, r)
	}

	return len(res)
}

func OptimisedReact(src string) int {
	shortest := len(src)

	for r := 'A'; r < 'Z'; r++ {
		candidate := removeUnit(src, r)
		stable := React(candidate)
		if stable < shortest {
			shortest = stable
		}
	}

	return shortest
}

func ReOptimisedReact(src string) int {
	units := detectUnits(src)

	shortest := len(src)

	for r := range units {
		candidate := removeUnit(src, r)
		stable := React(candidate)
		if stable < shortest {
			shortest = stable
		}
	}

	return shortest
}

func detectUnits(src string) (res map[rune]struct{}) {
	res = make(map[rune]struct{})
	for r := 'A'; r < 'Z'; r++ {
		if strings.ContainsAny(src, string([]rune{r, r + charDiff})) {
			res[r] = struct{}{}
		}
	}

	return
}

func removeUnit(src string, r rune) string {
	return strings.Replace(strings.Replace(src, string(r), "", -1), string(r+charDiff), "", -1)
}

func ReadInput() string {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	input, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	return string(input)
}

func main() {
	input := ReadInput()
	start := time.Now()
	stable := React(input)
	duration := time.Since(start)
	fmt.Printf("Got %d in %v\n", stable, duration)

	start = time.Now()
	shortest := OptimisedReact(input)
	duration = time.Since(start)
	fmt.Printf("Got %d in %v\n", shortest, duration)
}
