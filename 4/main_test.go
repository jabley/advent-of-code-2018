package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	input := `
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up
`

	sleepAnalysis := AnalyseShifts(input)

	if sleepAnalysis.SleepiestGuard() != 10 {
		t.Errorf("Could not find sleepliest guard")
	}

	if sleepAnalysis.MostAsleepMinute(10) != 24 {
		t.Errorf("Expected: %v but was %v", 24, sleepAnalysis.MostAsleepMinute(10))
	}

	guard, minute := sleepAnalysis.FrequentSleeper()

	if guard != 99 {
		t.Errorf("Expected frequent sleeper to be %s but was %d", "99", guard)
	}

	if minute != 45 {
		t.Errorf("Expected frequent minute to be %s but was %d", "45", minute)
	}
}
