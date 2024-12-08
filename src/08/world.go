package main

import "bufio"

type World struct {
	width  int
	height int
	nodes  map[rune][][2]int
}

func loadWorld(scanner *bufio.Scanner) *World {
	world := World{}
	world.nodes = make(map[rune][][2]int)

	var i int
	var line string
	for i = 0; scanner.Scan(); i++ {
		line = scanner.Text()
		for j, r := range line {
			if r != '.' {
				world.nodes[r] = append(world.nodes[r], [2]int{j, i})
			}
		}
	}
	world.width = len(line)
	world.height = i

	return &world
}

func (w *World) isPointOutOfBounds(x, y int) bool {
	return x < 0 || x >= w.width || y < 0 || y >= w.height
}

func (w *World) findAntinodes() map[[2]int]bool {
	antinodes := make(map[[2]int]bool)
	for _, channel := range w.nodes {
		for _, node := range channel {
			for _, otherNode := range channel {
				if node == otherNode {
					continue
				}

				xPos := node[0] + (otherNode[0]-node[0])*2
				yPos := node[1] + (otherNode[1]-node[1])*2
				if !w.isPointOutOfBounds(xPos, yPos) {
					antinodes[[2]int{xPos, yPos}] = true
				}

			}
		}
	}

	return antinodes
}
