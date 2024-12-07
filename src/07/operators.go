package main

import (
	"math"
	"strconv"
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

func Concat(x, y int) int {
	return parseInt(strconv.Itoa(x) + strconv.Itoa(y))
}
