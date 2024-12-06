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

func partOne() {
	file, scanner := readInput("input.txt")
	defer file.Close()

	world := loadWorld(scanner)
	guard := world.guard
	for !guard.IsOutOfBounds(world) {
		if !guard.MoveForward(world) {
			guard.TurnRight()
		}
	}

	fmt.Println("Part one:", len(guard.placesVisited)-1)
}
