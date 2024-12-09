package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) string {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	content := scanner.Text()

	return content
}

func fillSlice(s []int, i, repeat, val int) {
	for j := 0; j < repeat; j++ {
		s[i+j] = val
	}
}

func createDisk(content string) []int {
	size := 0
	for _, r := range content {
		size += int(r - '0')
	}
	return make([]int, size)
}

func populateDisk(disk []int, content string) {
	free := false
	i := 0
	id := 0
	for _, r := range content {
		repeat := int(r - '0')
		if free {
			fillSlice(disk, i, repeat, -1)
		} else if repeat != 0 {
			fillSlice(disk, i, repeat, id)
			id += 1
		}
		free = !free
		i += repeat
	}
}

func compactDisk(disk []int) {
	length := len(disk)
	i := 0
	j := length - 1
	for true {
		for disk[i] != -1 {
			i++
		}
		for disk[j] == -1 {
			j--
		}

		if i >= j {
			break
		}

		disk[i] = disk[j]
		disk[j] = -1
	}
}

func calculateChecksum(disk []int) int {
	checksum := 0
	for i, x := range disk {
		if disk[i] == -1 {
			continue
		}

		checksum += i * x
	}
	return checksum
}

func partOne() {
	content := readInput("input.txt")
	disk := createDisk(content)
	populateDisk(disk, content)
	compactDisk(disk)
	checksum := calculateChecksum(disk)
	fmt.Println("Part one:", checksum)
}
