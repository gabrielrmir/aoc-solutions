package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type wordSearch struct {
	word      string
	direction int
	column    int
	pos       int
}

func newWordSearch(word string, direction int, column int) *wordSearch {
	w := wordSearch{
		word:      word,
		direction: direction,
		column:    column,
		pos:       0,
	}
	return &w
}

func continueWord(w *wordSearch, line string) bool {
	if line[w.column+w.direction] != w.word[w.pos+1] {
		return false
	}
	w.pos += 1
	w.column += w.direction
	return true
}

func isWordComplete(w wordSearch) bool {
	return w.pos == len(w.word)-1
}

func readInput(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func searchNewWord(words []wordSearch, word string, line string) []wordSearch {
	newWords := words
	for i := 0; i < len(line); i++ {
		if word[0] != line[i] {
			continue
		}
		newWords = append(newWords, *newWordSearch(word, 0, i))
		if i <= len(line)-len(word) {
			newWords = append(newWords, *newWordSearch(word, 1, i))
		}
		if i >= len(word)-1 {
			newWords = append(newWords, *newWordSearch(word, -1, i))
		}
	}
	return newWords
}

func scanLine(words []wordSearch, line string) []wordSearch {
	var newWords []wordSearch
	for _, w := range words {
		if isWordComplete(w) {
			continue
		}

		if continueWord(&w, line) {
			newWords = append(newWords, w)
		}
	}

	newWords = searchNewWord(newWords, "XMAS", line)
	newWords = searchNewWord(newWords, "SAMX", line)

	return newWords
}

func countWords(scanner *bufio.Scanner) int {
	found := 0
	var words []wordSearch
	for scanner.Scan() {
		line := scanner.Text()
		found += strings.Count(line, "XMAS")
		found += strings.Count(line, "SAMX")
		words = scanLine(words, line)
		for _, w := range words {
			if w.pos == len(w.word)-1 {
				found += 1
			}
		}
	}
	return found
}

func partOne() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	found := countWords(scanner)
	fmt.Println("Part one:", found)
}
