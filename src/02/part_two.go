package main

import (
	"bufio"
	"fmt"
	"strings"
)

func removeAt(slice []int, at int) []int {
	var newSlice []int
	for i, x := range slice {
		if i != at {
			newSlice = append(newSlice, x)
		}
	}
	return newSlice
}

func countSafeWithTolerance(s *bufio.Scanner) int {
	safeTotal := 0

	for s.Scan() {
		line := s.Text()
		levels := parseInts(strings.Split(line, " "))

		if checkSafe(levels) {
			safeTotal += 1
			continue
		}

		for i := range levels {
			newLevels := removeAt(levels, i)
			if checkSafe(newLevels) {
				safeTotal += 1
				break
			}
		}
	}

	return safeTotal
}

func partTwo() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	safeTotal := countSafeWithTolerance(scanner)
	fmt.Println("Part two:", safeTotal)
}
