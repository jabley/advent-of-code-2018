package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

// A map of nodes with the set of child tasks
type dag map[rune]map[rune]struct{}

type runeSlice []rune

func (rs runeSlice) Len() int           { return len(rs) }
func (rs runeSlice) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
func (rs runeSlice) Less(i, j int) bool { return rs[i] < rs[j] }

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	order := CalculateOrder(f)
	fmt.Printf("%s\n", order)
}

func CalculateOrder(r io.Reader) string {
	graph := parseInput(r)

	steps := []rune{}

	for len(graph) > 0 {
		available := []rune{}

		for node, prereqs := range graph {
			if len(prereqs) == 0 {
				available = append(available, node)
			}
		}

		// If more than one step is ready, choose the step which is first alphabetically
		sort.Sort(runeSlice(available))

		steps = append(steps, available[0])

		for node, prereqs := range graph {
			for prereq := range prereqs {
				if prereq == available[0] {
					delete(graph[node], prereq)
				}
			}
		}

		delete(graph, available[0])
	}

	return string(steps)
}

func parseInput(r io.Reader) dag {
	graph := make(dag)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		prereq := rune(line[5])
		step := rune(line[36])

		if _, ok := graph[step]; !ok {
			graph[step] = make(map[rune]struct{})
		}
		if _, ok := graph[prereq]; !ok {
			graph[prereq] = make(map[rune]struct{})
		}

		graph[step][prereq] = struct{}{}
	}

	return graph
}
