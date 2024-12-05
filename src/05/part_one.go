package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

func readInput(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func readRules(scanner *bufio.Scanner) map[string][]string {
	rules := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rule := strings.Split(line, "|")
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}
	return rules
}

func checkUpdate(update []string, rules map[string][]string) bool {
	for i, page := range update {
		for _, pageAfter := range update[i+1:] {
			if slices.Contains(rules[pageAfter], page) {
				return false
			}
		}
	}
	return true
}

func partOne() {
	file, scanner := readInput("input.txt")
	defer file.Close()

	rules := readRules(scanner)
	sum := 0
	for scanner.Scan() {
		update := strings.Split(scanner.Text(), ",")
		if checkUpdate(update, rules) {
			sum += parseInt(update[len(update)/2])
		}
	}

	fmt.Println("Part one:", sum)
}
