package main

import "bufio"

type World struct {
	width     int
	height    int
	obstacles map[[2]int]bool
	guard     *Guard
}

func (w *World) HasObstacle(x, y int) bool {
	obstacle, ok := w.obstacles[[2]int{x, y}]
	return ok && obstacle
}

func loadWorld(scanner *bufio.Scanner) *World {
	world := World{}
	world.obstacles = make(map[[2]int]bool)

	var (
		i    int
		line string
	)

	for i = 0; scanner.Scan(); i++ {
		line = scanner.Text()
		for j, c := range line {
			if c == '#' {
				world.obstacles[[2]int{j, i}] = true
			} else if c != '.' {
				guard := Guard{
					x:      j,
					y:      i,
					facing: directionRune[c],
				}
				guard.placesVisited = make(map[[2]int]bool)
				guard.placesVisited[[2]int{guard.x, guard.y}] = true
				world.guard = &guard
			}
		}
	}

	world.width = len(line)
	world.height = i

	return &world
}
