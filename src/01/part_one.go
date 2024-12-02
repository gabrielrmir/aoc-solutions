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
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readLists(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids := strings.Split(scanner.Text(), "   ")
		left = append(left, parseInt(ids[0]))
		right = append(right, parseInt(ids[len(ids)-1]))
	}

	return left, right
}

func totalDistance(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += abs(left[i] - right[i])
	}

	return sum
}

func main() {
	left, right := readLists("input.txt")
	dist := totalDistance(left, right)
	fmt.Println(dist)
}
