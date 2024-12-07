package main

import "fmt"

func partTwo() {
	fns := []operation{Sum, Mult, Concat}
	fmt.Println("Part two:", evalFile("input.txt", fns))
}
