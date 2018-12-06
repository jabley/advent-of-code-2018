package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const charDiff = 'a' - 'A'

func React(src string) string {
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

	return string(res)
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	input, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	stable := React(string(input))

	fmt.Println(len(stable))
}
