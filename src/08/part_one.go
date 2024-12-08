package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func partOne(world *World) {
	antinodes := world.findAntinodes()
	fmt.Println("Part one:", len(antinodes))
}
