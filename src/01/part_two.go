package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readListsMap(filename string) (map[string]int, map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	left := make(map[string]int)
	right := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids := strings.Split(scanner.Text(), "   ")
		left[ids[0]] = left[ids[0]] + 1
		right[ids[1]] = right[ids[1]] + 1
	}

	return left, right
}

func totalSimilarity(left map[string]int, right map[string]int) int {
	similarity := 0
	for key := range left {
		similarity += parseInt(key) * right[key]
	}
	return similarity
}

func partTwo() {
	left, right := readListsMap("input.txt")
	similarity := totalSimilarity(left, right)
	fmt.Println("Part two:", similarity)
}
