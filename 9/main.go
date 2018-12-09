package main

import (
	"container/ring"
)

func Play(players, marbles int) int {
	circle := &ring.Ring{Value: 0}
	scores := make([]int, players)

	maxScore := 0

	for i := 1; i <= marbles; i++ {
		if i%23 == 0 {
			player := i % players
			circle = circle.Move(-8)
			scores[player] += i + circle.Unlink(1).Value.(int)
			if maxScore < scores[player] {
				maxScore = scores[player]
			}
		} else {
			circle = circle.Move(1)
			circle.Link(&ring.Ring{Value: i})
		}
		circle = circle.Move(1)
	}

	return maxScore
}

func main() {

}
