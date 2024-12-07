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

func readInput(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return file, scanner
}

func parseInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return x
}

func readLine(line string) (int, []int) {
	split := strings.Split(line, ": ")
	snums := strings.Split(split[1], " ")
	nums := make([]int, len(snums))
	for i, s := range snums {
		nums[i] = parseInt(s)
	}
	return parseInt(split[0]), nums
}

func toBase(x, base int) []int {
	digits := make([]int, 0)
	for x >= base {
		digits = append(digits, x%base)
		x = x / base
	}
	digits = append(digits, x%base)
	slices.Reverse(digits)
	return digits
}

func checkPossible(result int, nums []int, fns []operation) bool {
	possibleArragements := Pow(len(fns), len(nums)-1)
	for i := 0; i < possibleArragements; i++ {
		acc := nums[0]
		opSequence := toBase(i+possibleArragements, len(fns))
		for j, x := range nums[1:] {
			acc = fns[opSequence[j+1]](acc, x)
		}
		if acc == result {
			return true
		}
	}
	return false
}

func evalFile(filename string, fns []operation) int {
	file, scanner := readInput(filename)
	defer file.Close()

	sum := 0
	for scanner.Scan() {
		result, nums := readLine(scanner.Text())
		if checkPossible(result, nums, fns) {
			sum += result
		}
	}

	return sum
}

func partOne() {
	fns := []operation{Sum, Mult}
	fmt.Println("Part one:", evalFile("input.txt", fns))
}
