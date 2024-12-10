package main

import "fmt"

func (w *World) findTrailheadUniqueScore(x, y int) int {
	if w.at(x, y) == 9 {
		return 1
	}

	checks := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	score := 0
	for _, check := range checks {
		offX := x + check[0]
		offY := y + check[1]
		if w.at(offX, offY) == w.world[y][x]+1 {
			score += w.findTrailheadUniqueScore(offX, offY)
		}
	}
	return score
}

func partTwo(world *World) {
	score := 0
	for trail := range world.trailheads {
		score += world.findTrailheadUniqueScore(trail[0], trail[1])
	}
	fmt.Println("Part two:", score)
}
