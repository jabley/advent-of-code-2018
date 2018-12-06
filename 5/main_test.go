package main

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{"dabAcCaCBAcCcaDA", 10},
		{"aA", 0},
		{"abBA", 0},
		{"abAB", 4},
		{"aabAAB", 6},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			stable := React(tt.input)

			if stable != tt.expectedLen {
				t.Errorf("Expected %d but got %d", tt.expectedLen, stable)
			}
		})
	}
}

func TestFoo(t *testing.T) {
	// fmt.Printf("a: %d, A: %d \n", 'a', 'A')
}

func TestOptimisedReactExamples(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{"dabAcCaCBAcCcaDA", 4},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			stable := OptimisedReact(tt.input)

			if stable != tt.expectedLen {
				t.Errorf("Expected %d but got %d", tt.expectedLen, stable)
			}
		})
	}
}

func BenchmarkReact(b *testing.B) {
	input := ReadInput()

	for i := 0; i < b.N; i++ {
		React(input)
	}
}

func BenchmarkOptimisedReact(b *testing.B) {
	input := ReadInput()

	for i := 0; i < b.N; i++ {
		OptimisedReact(input)
	}
}

func BenchmarkReOptimisedReact(b *testing.B) {
	input := ReadInput()

	for i := 0; i < b.N; i++ {
		ReOptimisedReact(input)
	}
}
