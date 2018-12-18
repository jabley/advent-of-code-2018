package main

import (
	"fmt"
	"strconv"
	"time"
)

type process struct {
	scores    []byte
	firstElf  int
	secondElf int
}

func new(scores []byte, firstElf, secondElf int) *process {
	return &process{
		scores:    scores,
		firstElf:  firstElf,
		secondElf: secondElf,
	}
}

func (p *process) followingScores(nRecipes int) string {
	return p.scoreboard()[nRecipes : nRecipes+10]
}

func (p *process) mix() {
	score := []byte(strconv.Itoa(recipeScoreToInt(p.scores[p.firstElf]) + recipeScoreToInt(p.scores[p.secondElf])))
	p.scores = append(p.scores, score...)

	p.firstElf = p.pickNewRecipe(p.firstElf)
	p.secondElf = p.pickNewRecipe(p.secondElf)
}

func (p *process) pickNewRecipe(elf int) int {
	return (elf + 1 + recipeScoreToInt(p.scores[elf])) % len(p.scores)
}

func recipeScoreToInt(score byte) int {
	return int(score - '0')
}

func (p *process) scoreboard() string {
	return string(p.scores)
}

const input = 793061

func main() {
	start := time.Now()

	p := new([]byte{'3', '7'}, 0, 1)

	for len(p.scores) < input+10 {
		p.mix()
	}

	fmt.Printf("Took %v\n", time.Since(start))
	fmt.Printf("part 1 %s\n", p.followingScores(input))
}
