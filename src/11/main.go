package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func countDigits(x int) int {
	return len(strconv.Itoa(x))
}

func parseInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

var memo map[[2]int]int

func main() {
	content := readInput("input.txt")
	split := strings.Split(content, " ")

	rocks := make([]int, 0, len(split))
	for _, s := range split {
		rocks = append(rocks, parseInt(s))
	}

	memo = make(map[[2]int]int)

	fmt.Println("Part one:", BlinkRocks(rocks, 25))
	fmt.Println("Part two:", BlinkRocks(rocks, 75))
}
