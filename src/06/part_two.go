package main

import "fmt"

type Position struct {
	x      int
	y      int
	facing Direction
}

func (g *Guard) SetPosition(pos Position) {
	g.x = pos.x
	g.y = pos.y
	g.facing = pos.facing
	g.placesVisited[[2]int{g.x, g.y}] = true
}

func (g *Guard) GetPosition() Position {
	return Position{
		x:      g.x,
		y:      g.y,
		facing: g.facing,
	}
}

func partTwo() {
	file, scanner := readInput("input.txt")
	defer file.Close()

	world := loadWorld(scanner)
	guard := world.guard
	start := guard.GetPosition()

	for !guard.IsOutOfBounds(world) {
		if !guard.MoveForward(world) {
			guard.TurnRight()
		}
	}

	delete(guard.placesVisited, [2]int{start.x, start.y})
	delete(guard.placesVisited, [2]int{guard.x, guard.y})

	possibleObstacles := make([][2]int, len(guard.placesVisited))
	i := 0
	for key := range guard.placesVisited {
		possibleObstacles[i] = key
		i++
	}

	loopCount := 0
	for _, newObstacle := range possibleObstacles {
		guard.placesVisited = make(map[[2]int]bool)
		guard.SetPosition(start)
		world.obstacles[newObstacle] = true
		isLoop := false
		turns := make(map[Position]bool)

		for !guard.IsOutOfBounds(world) {
			if !guard.MoveForward(world) {
				pos := guard.GetPosition()
				if _, ok := turns[pos]; ok {
					isLoop = true
					break
				}
				turns[pos] = true
				guard.TurnRight()
			}
		}
		delete(world.obstacles, newObstacle)
		if isLoop {
			loopCount += 1
		}
	}

	fmt.Println("Part two:", loopCount)
}
