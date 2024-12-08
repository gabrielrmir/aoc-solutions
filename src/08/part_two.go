package main

import "fmt"

func partTwo(world *World) {
	antinodes := world.findResonantAntinodes()
	fmt.Println("Part two:", len(antinodes))
}
