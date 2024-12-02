package main

import (
	"log"
	"strconv"
)

func parseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
