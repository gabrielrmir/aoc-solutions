package main

import (
	"log"
	"strconv"
)

func parseInts(s []string) []int {
	ints := make([]int, len(s))
	for i, x := range s {
		num, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = num
	}
	return ints
}
