package main

func main() {
	file, scanner := readInput("input.txt")
	defer file.Close()
	world := loadWorld(scanner)

	partOne(world)
	partTwo(world)
}
