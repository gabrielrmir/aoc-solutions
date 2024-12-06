package main

import "fmt"

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
	possibleObstacles := guard.placesVisited

	loopCount := 0
	for newObstacle := range possibleObstacles {
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
