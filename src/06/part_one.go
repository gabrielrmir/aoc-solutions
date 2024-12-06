package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction int

const (
	DirectionRight Direction = iota
	DirectionDown
	DirectionLeft
	DirectionUp
)

var directionRune = map[rune]Direction{
	'>': DirectionRight,
	'v': DirectionDown,
	'<': DirectionLeft,
	'^': DirectionUp,
}

type Guard struct {
	x             int
	y             int
	facing        Direction
	placesVisited map[[2]int]bool
}

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

func (g *Guard) IsOutOfBounds(world *World) bool {
	return g.x < 0 || g.y < 0 || g.x >= world.width || g.y >= world.height
}

func (g *Guard) TurnRight() {
	g.facing = (g.facing + 1) % 4
}

func (g *Guard) MoveForward(world *World) bool {
	newX := g.x
	newY := g.y
	if g.facing == DirectionRight {
		newX += 1
	} else if g.facing == DirectionDown {
		newY += 1
	} else if g.facing == DirectionLeft {
		newX -= 1
	} else if g.facing == DirectionUp {
		newY -= 1
	}
	if world.HasObstacle(newX, newY) {
		return false
	}
	g.x = newX
	g.y = newY
	g.placesVisited[[2]int{g.x, g.y}] = true
	return true
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

func readInput(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func partOne() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	world := loadWorld(scanner)
	for !world.guard.IsOutOfBounds(world) {
		moved := world.guard.MoveForward(world)
		if !moved {
			world.guard.TurnRight()
		}
	}
	fmt.Println("Part one:", len(world.guard.placesVisited)-1)
}
