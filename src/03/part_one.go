package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func readInput(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func scanMul(scanner *bufio.Scanner) int {
	sum := 0
	r, _ := regexp.Compile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindAllStringSubmatch(line, -1)
		for _, s := range match {
			sum += parseInt(s[1]) * parseInt(s[2])
		}
	}
	return sum
}

func partOne() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	sum := scanMul(scanner)
	fmt.Println("Part one:", sum)
}
