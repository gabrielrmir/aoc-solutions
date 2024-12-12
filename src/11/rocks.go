package main

import (
	"strconv"
)

func Blink(rock int, depth int) int {
	if l, ok := memo[[2]int{rock, depth}]; ok {
		return l
	}

	isEven := countDigits(rock)%2 == 0

	if depth <= 1 {
		if isEven {
			return 2
		}
		return 1
	}

	if rock == 0 {
		memo[[2]int{rock, depth}] = Blink(1, depth-1)
		return memo[[2]int{rock, depth}]
	} else if isEven {
		s := strconv.Itoa(rock)
		left := parseInt(s[:len(s)/2])
		right := parseInt(s[len(s)/2:])

		memo[[2]int{rock, depth}] = Blink(left, depth-1) + Blink(right, depth-1)
		return memo[[2]int{rock, depth}]
	} else {
		memo[[2]int{rock, depth}] = Blink(rock*2024, depth-1)
		return memo[[2]int{rock, depth}]
	}
}

func BlinkRocks(rocks []int, n int) int {
	numRocks := 0
	for _, rock := range rocks {
		numRocks += Blink(rock, n)
	}
	return numRocks
}
