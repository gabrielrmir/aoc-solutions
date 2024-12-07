package main

import (
	"math"
)

type operation func(int, int) int

func Sum(x, y int) int {
	return x + y
}

func Mult(x, y int) int {
	return x * y
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
