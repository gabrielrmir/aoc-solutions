package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInts(s []string) []int {
	ints := make([]int, len(s))
	for i, x := range s {
		num, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = num
	}
	return ints
}

func readInput(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func countSafe(s *bufio.Scanner) int {
	safeTotal := 0

	for s.Scan() {
		line := s.Text()
		levels := parseInts(strings.Split(line, " "))

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

		if isSafe {
			safeTotal += 1
		}
	}

	return safeTotal
}

func main() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	safeTotal := countSafe(scanner)
	fmt.Println(safeTotal)
}
