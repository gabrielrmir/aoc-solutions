package main

import (
	"bufio"
)

type World struct {
	world  [][]rune
	width  int
	height int
}

func NewWorld(scanner *bufio.Scanner) *World {
	world := World{}

	var i int
	var line string
	for i = 0; scanner.Scan(); i++ {
		line = scanner.Text()
		row := make([]rune, len(line))
		for j, r := range line {
			row[j] = r
		}
		world.world = append(world.world, row)
	}

	world.width = len(line)
	world.height = i

	return &world
}

func (w *World) At(x, y int) rune {
	if x < 0 || y < 0 || x >= w.width || y >= w.height {
		return 0
	}
	return w.world[y][x]
}

func (w *World) ExploreRegion(plots map[[2]int]int, explored [][]bool, x, y int) {
	at := w.At(x, y)
	explored[y][x] = true
	plots[[2]int{x, y}] = 0
	offsets := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, offset := range offsets {
		xOff := x + offset[0]
		yOff := y + offset[1]
		atOff := w.At(xOff, yOff)
		if atOff != 0 && at == atOff && !explored[yOff][xOff] {
			w.ExploreRegion(plots, explored, xOff, yOff)
		}
		if at != atOff {
			plots[[2]int{x, y}] += 1
		}
	}
}

func (w *World) GetRegions() []Region {
	regions := make([]Region, 0)
	explored := make([][]bool, w.height)
	for i := range explored {
		explored[i] = make([]bool, w.width)
	}

	for y, row := range w.world {
		for x := range row {
			if !explored[y][x] {
				plots := make(map[[2]int]int)
				w.ExploreRegion(plots, explored, x, y)
				regions = append(regions, *NewRegion(w.At(x, y), plots))
			}
		}
	}

	return regions
}
