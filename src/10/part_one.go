package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type World struct {
	world      [][]int
	width      int
	height     int
	trailheads map[[2]int]bool
}

func (w *World) at(x, y int) int {
	if x < 0 || y < 0 || x >= w.width || y >= w.height {
		return -1
	}
	return w.world[y][x]
}

func (w *World) findTrailheadTops(tops map[[2]int]bool, x, y int) {
	if w.at(x, y) == 9 {
		tops[[2]int{x, y}] = true
		return
	}

	checks := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, check := range checks {
		offX := x + check[0]
		offY := y + check[1]
		if w.at(offX, offY) == w.world[y][x]+1 {
			w.findTrailheadTops(tops, offX, offY)
		}
	}
}

func (w *World) findTrailheadScore(x, y int) int {
	tops := make(map[[2]int]bool)
	w.findTrailheadTops(tops, x, y)
	return len(tops)
}

func loadWorld(scanner *bufio.Scanner) *World {
	world := World{}
	world.trailheads = make(map[[2]int]bool)
	i := 0
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		l := make([]int, len(line))
		for j, r := range line {
			if r == '0' {
				world.trailheads[[2]int{j, i}] = true
			}
			l[j] = int(r - '0')
		}
		world.world = append(world.world, l)
		i++
	}
	world.height = i
	world.width = len(line)

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

func partOne(world *World) {
	score := 0
	for trail := range world.trailheads {
		score += world.findTrailheadScore(trail[0], trail[1])
	}
	fmt.Println("Part one:", score)
}

