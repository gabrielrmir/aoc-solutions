package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func checkSafe(levels []int) bool {
	isSafe := true
	dir := sign(levels[1] - levels[0])

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if sign(diff) != dir {
			isSafe = false
			break
		}

		if diff == 0 || abs(diff) > 3 {
			isSafe = false
			break
		}
	}

	return isSafe
}

func countSafe(s *bufio.Scanner) int {
	safeTotal := 0
	for s.Scan() {
		line := s.Text()
		levels := parseInts(strings.Split(line, " "))
		if checkSafe(levels) {
			safeTotal += 1
		}
	}
	return safeTotal
}

func partOne() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	safeTotal := countSafe(scanner)
	fmt.Println(safeTotal)
}
