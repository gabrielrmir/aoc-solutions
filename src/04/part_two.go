package main

import (
	"bufio"
	"fmt"
)

func getWord(r rune) string {
	if r == 'M' {
		return "MAS"
	} else if r == 'S' {
		return "SAM"
	}
	return ""
}

func countX(scanner *bufio.Scanner) int {
	found := 0
	var wordPairs [][2]wordSearch
	for scanner.Scan() {
		line := scanner.Text()

		var newWordPairs [][2]wordSearch
		for _, wordPair := range wordPairs {
			if !continueWord(&wordPair[0], line) || !continueWord(&wordPair[1], line) {
				continue
			}
			if isWordComplete(wordPair[0]) && isWordComplete(wordPair[1]) {
				found += 1
				continue
			} else if isWordComplete(wordPair[0]) || isWordComplete(wordPair[1]) {
				continue
			}
			newWordPairs = append(newWordPairs, wordPair)
		}
		wordPairs = newWordPairs

		for i, r := range line {
			if i >= len(line)-2 {
				continue
			}

			w1 := getWord(r)
			w2 := getWord(rune(line[i+2]))
			if (w1 != "") && (w2 != "") {
				var newWordPair [2]wordSearch
				newWordPair[0] = *newWordSearch(w1, 1, i)
				newWordPair[1] = *newWordSearch(w2, -1, i+2)
				wordPairs = append(wordPairs, newWordPair)
			}
		}
	}
	return found
}

func partTwo() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	found := countX(scanner)
	fmt.Println("Part two:", found)
}
