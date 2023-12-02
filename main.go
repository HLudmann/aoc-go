package main

import (
	"fmt"
	a23d1 "hludmann/aoc/2023/day1"
	a23d2 "hludmann/aoc/2023/day2"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1_input, err := os.ReadFile("inputs/2023/day1.txt")
	check(err)
	d1p1 := a23d1.Puzzle1(string(d1_input))
	d1p2 := a23d1.Puzzle2(string(d1_input))
	fmt.Printf("Day 1\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", d1p1, d1p2)

	d2_input, err := os.ReadFile("inputs/2023/day2.txt")
	check(err)
	d2p1 := a23d2.Puzzle1(string(d2_input))
	d2p2 := a23d2.Puzzle2(string(d2_input))
	fmt.Printf("Day 2\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", d2p1, d2p2)
}
