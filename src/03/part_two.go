package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func scanExtended(scanner *bufio.Scanner) int {
	sum := 0
	enabled := true
	r, _ := regexp.Compile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)|do\\(\\)|don't\\(\\)")
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindAllStringSubmatch(line, -1)
		for _, s := range match {
			if enabled && strings.HasPrefix(s[0], "mul(") {
				sum += parseInt(s[1]) * parseInt(s[2])
			} else if strings.HasPrefix(s[0], "don't(") {
				enabled = false
			} else if strings.HasPrefix(s[0], "do(") {
				enabled = true
			}
		}
	}
	return sum
}

func partTwo() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	sum := scanExtended(scanner)
	fmt.Println("Part two:", sum)
}
