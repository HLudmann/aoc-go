package y2023

import (
	"fmt"
	"os"
	"strings"
)

func posMod(x, y int) int {
	if m := x % y; m < 0 {
		return y + m
	}
	return x % y
}

func inputToGardensSizeAndStart(input string) (map[Pos]int, int, Pos) {
	gardens := make(map[Pos]int)
	var s Pos
	var size int
	for i, line := range toLines(input) {
		for j, val := range line {
			if val == '.' || val == 'S' {
				gardens[Pos{i, j}] = 1
			}
			if val == 'S' {
				s = Pos{i, j}
			}
		}
		size++
	}

	return gardens, size, s
}

func countReachable(gardens map[Pos]int, size int, start Pos, steps int) int {
	dist := map[Pos]int{start: 0}
	queue := []Pos{start}
	shifts := []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) != 0 {
		var nq []Pos
		for _, p := range queue {
			for _, shift := range shifts {
				n := p.Add(shift)
				newD := dist[p] + 1
				oldD := dist[n]
				if oldD == 0 && n != start {
					oldD = steps + 1
				}
				if gardens[Pos{posMod(n.x, size), posMod(n.y, size)}] == 1 && newD < oldD {
					dist[n] = newD
					nq = append(nq, n)
				}
			}
		}
		queue = nq
	}

	var count int
	comp := steps % 2
	for _, val := range dist {
		if val%2 == comp {
			count++
		}
	}
	return count
}

func day21Part1(input string) string {
	gardens, size, s := inputToGardensSizeAndStart(input)

	return fmt.Sprint(countReachable(gardens, size, s, 64))
}

func day21Part2(input string) string {
	gardens, size, s := inputToGardensSizeAndStart(input)
	steps := 26501365

	h := countReachable(gardens, size, s, size/2)
	hs := countReachable(gardens, size, s, size+size/2)
	h2s := countReachable(gardens, size, s, 2*size+size/2)

	a := (h2s + h - 2*hs) / 2
	b := hs - h - a
	c := h
	n := (steps - s.x) / size

	count := a*n*n + b*n + c
	return fmt.Sprint(count)
}

func Day21(test bool) {
	path := "inputs/2023/day21.txt"
	if test {
		path = strings.Replace(path, "day21", "day21-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day21Part1(string(input))
	p2 := day21Part2(string(input))

	fmt.Printf("Day 21\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
