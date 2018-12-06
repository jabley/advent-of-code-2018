package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type nap struct {
	start time.Time
	end   time.Time
}

type guard struct {
	id   int
	naps []nap
}

func totalSleep(naps []nap) int {
	res := 0

	for _, n := range naps {
		res += int(n.end.Sub(n.start).Minutes())
	}

	return res
}

func mostAsleepMinute(naps []nap) (int, int) {
	res := -1

	minutes := make([]int, 60)
	for _, n := range naps {
		for i := n.start.Minute(); i < n.end.Minute(); i++ {
			minutes[i]++
		}
	}

	max := -1
	for i, j := range minutes {
		if j > max {
			max = j
			res = i
		}
	}

	return res, max
}

type SleepAnalysis struct {
	guards map[int][]nap
}

func (sa *SleepAnalysis) SleepiestGuard() (res int) {
	max := 0

	for g, naps := range sa.guards {
		sleep := totalSleep(naps)

		// fmt.Printf("Guard %s slept for %d total minutes\n", g, sleep)

		if sleep > max {
			max = sleep
			res = g
		}
	}

	return
}

func (sa *SleepAnalysis) MostAsleepMinute(guardID int) int {
	minutes := make([]int, 59)

	for _, n := range sa.guards[guardID] {
		for i := n.start.Minute(); i < n.end.Minute(); i++ {
			minutes[i]++
		}
	}

	max := 0
	res := -1

	for i, j := range minutes {
		if j > max {
			res = i
			max = j
		}
	}

	return res
}

func (sa *SleepAnalysis) FrequentSleeper() (g int, m int) {
	mostTimes := 0

	for guard := range sa.guards {
		minute, times := mostAsleepMinute(sa.guards[guard])

		if times > mostTimes {
			g = guard
			m = minute
			mostTimes = times
		}
	}

	return
}

func AnalyseShifts(input string) *SleepAnalysis {
	return &SleepAnalysis{guards: parseInput(input)}
}

func parseInput(input string) map[int][]nap {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var guardID int

	guards := make(map[int][]nap)

	n := 0
	for scanner.Scan() {
		n++
		line := scanner.Text()

		// println(fmt.Sprintf("(Doing line <%s>", line))

		if strings.HasSuffix(line, " begins shift") {
			guardID = extractGuardID(line)
			if _, ok := guards[guardID]; !ok {
				guards[guardID] = make([]nap, 0)
			}
		} else if strings.HasSuffix(line, " falls asleep") {
			start := parseTime(line)
			guards[guardID] = append(guards[guardID], nap{start: start})
		} else if strings.HasSuffix(line, " wakes up") {
			if guards[guardID] == nil {
				panic(fmt.Sprintf("not fallen asleep yet! at line %d", n))
			}
			end := parseTime(line)
			guards[guardID][len(guards[guardID])-1].end = end
		} else {
			// println(fmt.Sprintf("Unhandled line: <%s>", line))
		}

	}

	return guards
}

func extractGuardID(line string) int {
	i := strings.Index(line, "#")
	s := line[i+1 : len(line)-len(" begins shift")]
	g, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse guard ID <%s> as int", s))
	}

	return int(g)
}

func parseTime(line string) time.Time {
	t, err := time.Parse("2006-01-02 15:04", line[1:17])
	if err != nil {
		panic("uable to parse time")
	}
	return t
}

func main() {
	lines, err := readLines("input.txt")

	if err != nil {
		panic(err)
	}

	sort.Strings(lines)

	analysis := AnalyseShifts(strings.Join(lines, "\n"))

	sleepiestGuard := analysis.SleepiestGuard()
	fmt.Printf("sleepiest guard: %d\n", sleepiestGuard)
	fmt.Printf("most snoozed minute: %d\n", analysis.MostAsleepMinute(sleepiestGuard))

	guard, minute := analysis.FrequentSleeper()

	fmt.Printf("frequent sleeper: #%d at %d = %d\n", guard, minute, guard*minute)

}

func readLines(file string) (lines []string, err error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		const delim = '\n'
		line, err := r.ReadString(delim)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	return lines, nil
}
