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

	// order := CalculateOrder(f)
	// fmt.Printf("%s\n", order)
	totalTime := CalculateTime(f, 5, 60)
	fmt.Printf("%d\n", totalTime)
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

func CalculateTime(r io.Reader, nWorkers int, fixedJobCost int) int {
	graph := parseInput(r)

	// workers is a map of step duration keyed by step
	workers := make(map[rune]int)
	totalTime := 0

	// While we have nodes to work on, or workers are still working
	for len(graph) > 0 || len(workers) > 0 {
		// are there any workers about to become free?
		for s, t := range workers {
			workers[s]--

			if t == 1 {
				// worker will finish on the next second
				// fmt.Printf("Worker is finishing step %c\n", s)

				delete(workers, s)
				delete(graph, s)

				for node := range graph {
					delete(graph[node], s)
				}
			}
		}

		available := []rune{}

		for node, prereqs := range graph {
			if len(prereqs) == 0 {
				available = append(available, node)
			}
		}

		// If more than one step is ready, choose the step which is first alphabetically
		sort.Sort(runeSlice(available))

		// fmt.Printf("We have the following steps available to work on: %s\n", string(available))

		// Are there empty workers to assign a job to?
		for len(workers) < nWorkers {
			if len(available) == 0 {
				break // need to let time tick on
			}

			currStep := available[0]
			if workers[currStep] == 0 {
				workers[currStep] = calculateStepTime(currStep, fixedJobCost)
				delete(graph, currStep)
				// fmt.Printf("Assigning %c to worker at time %d\n", currStep, totalTime)
				available = remove(available, currStep)
			}
		}

		// fmt.Printf("at time %v, we have workers %v\n", totalTime, workers)

		totalTime++
	}

	return totalTime - 1
}

func calculateStepTime(r rune, fixedJobCost int) int {
	return int(r-'A') + 1 + fixedJobCost
}

func remove(runes []rune, del rune) []rune {
	for i, r := range runes {
		if r == del {
			return append(runes[:i], runes[i+1:]...)
		}
	}
	return runes
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
