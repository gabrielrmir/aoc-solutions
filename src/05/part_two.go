package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func reorderUpdate(update []string, rules map[string][]string) []string {
	sort.Slice(update, func(i, j int) bool {
		return !slices.Contains(rules[update[j]], update[i])
	})
	return update
}

func partTwo() {
	file, scanner := readInput("input.txt")
	defer file.Close()

	rules := readRules(scanner)
	sum := 0
	for scanner.Scan() {
		update := strings.Split(scanner.Text(), ",")
		if !checkUpdate(update, rules) {
			newUpdate := reorderUpdate(update, rules)
			sum += parseInt(newUpdate[len(newUpdate)/2])
		}
	}

	fmt.Println("Part two:", sum)
}
