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

func calculatePrice(regions []Region) int {
	price := 0
	for _, region := range regions {
		price += region.area * region.perimeter
	}
	return price
}

func main() {
	file, scanner := readInput("input.txt")
	defer file.Close()

	world := NewWorld(scanner)
	regions := world.GetRegions()

	price := calculatePrice(regions)
	fmt.Println("Part one:", price)
}
