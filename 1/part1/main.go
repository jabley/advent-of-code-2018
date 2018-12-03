package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	var currentFrequency int64

	for scanner.Scan() {
		line := scanner.Text()
		change, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		currentFrequency += change
	}

	fmt.Println(currentFrequency)
}
