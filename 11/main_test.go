package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		serial   int
		expected string
	}{
		{18, "33,45"},
		{42, "21,61"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.serial), func(t *testing.T) {
			actual := BiggestClusterFor3(tt.serial)
			if actual.String() != tt.expected {
				t.Errorf("Expected %s but was %s", tt.expected, actual)
			}
		})
	}
}

func TestBuildSummedAreaTable(t *testing.T) {
	// Example from wikipedia image https://en.wikipedia.org/wiki/Summed-area_table at 2018-12-11
	cells := make([][]int, 6)

	cells[0] = []int{31, 12, 13, 24, 30, 1}
	cells[1] = []int{2, 26, 17, 23, 8, 35}
	cells[2] = []int{4, 9, 21, 15, 28, 34}
	cells[3] = []int{33, 10, 22, 16, 27, 3}
	cells[4] = []int{5, 29, 20, 14, 11, 32}
	cells[5] = []int{36, 25, 18, 19, 7, 6}

	expected := [][]int{
		[]int{31, 43, 56, 80, 110, 111},
		[]int{33, 71, 101, 148, 186, 222},
		[]int{37, 84, 135, 197, 263, 333},
		[]int{70, 127, 200, 278, 371, 444},
		[]int{75, 161, 254, 346, 450, 555},
		[]int{111, 222, 333, 444, 555, 666},
	}
	sumTable := BuildSumTable(cells)

	if !reflect.DeepEqual(sumTable, expected) {
		t.Errorf("Expected %v but got %v", expected, sumTable)
	}

	tests := []struct {
		x, y     int
		size     int
		expected int
	}{
		{0, 0, 2, 71},
		{0, 0, 3, 135},
		{0, 0, 6, 666},
		{0, 0, 5, 450},
		{1, 0, 5, 555 - 75},
		{0, 1, 5, 555 - 110},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d with size %d", tt.x, tt.y, tt.size), func(t *testing.T) {
			p := CalculateSquarePower(sumTable, tt.x, tt.y, tt.size)
			if p != tt.expected {
				t.Errorf("Expected square power to be %d but was %d", tt.expected, p)
			}
		})
	}
}
