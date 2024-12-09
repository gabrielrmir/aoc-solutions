package main

import "fmt"

func freeSpaceAt(disk []int, i int) int {
	size := 0
	for i < len(disk)-1 && disk[i] == -1 {
		size++
		i++
	}
	return size
}

func jumpBackwards(disk []int, j, id int) (int, int) {
	for disk[j] == id || disk[j] == -1 {
		j--
		if j <= 0 {
			return -1, -1
		}
	}

	size := 1
	newId := disk[j]
	for j > 0 && disk[j-1] == newId {
		j--
		size++
	}

	return j, size
}

func searchFreeSpace(disk []int, size int) int {
	var i int
	for i = 0; i < len(disk); i++ {
		if disk[i] != -1 {
			continue
		}
		if freeSpaceAt(disk, i) >= size {
			return i
		}
	}
	return -1
}

func compactDiskWhole(disk []int) {
	j := len(disk) - 1
	id := -1
	for {
		pos, size := jumpBackwards(disk, j, id)
		if pos < 0 {
			break
		}
		j = pos
		id = disk[j]

		i := searchFreeSpace(disk, size)
		if i < 0 || i > j {
			continue
		}

		fillSlice(disk, i, size, id)
		fillSlice(disk, j, size, -1)
	}
}

func partTwo() {
	content := readInput("input.txt")
	disk := createDisk(content)
	populateDisk(disk, content)
	compactDiskWhole(disk)
	checksum := calculateChecksum(disk)
	fmt.Println("Part two:", checksum)
}
