package y2023

import (
	"fmt"
	"os"
	"strings"
)

func galaxiesAfterExpension(universe []string, rate int) (galaxies []Pos) {
	rowsCounts := make([]int, len(universe))
	colsCounts := make([]int, len(universe[0]))

	for i, rows := range universe {
		for j, val := range rows {
			if val == '#' {
				rowsCounts[i] += 1
				colsCounts[j] += 1
			}
		}
	}

	var xShift, yShift int
	for x, row := range universe {
		yShift = 0
		for y, val := range row {
			if val == '#' {
				galaxies = append(galaxies, Pos{x + xShift, y + yShift})
			}
			if colsCounts[y] == 0 {
				yShift += (rate - 1)
			}
		}
		if rowsCounts[x] == 0 {
			xShift += (rate - 1)
		}
	}

	return
}

func day11Part1(input string) string {
	galaxies := galaxiesAfterExpension(toLines(input), 2)

	sum := 0
	for i, start := range galaxies[:len(galaxies)-1] {
		for _, end := range galaxies[i+1:] {
			sum += Abs(start.x-end.x) + Abs(start.y-end.y)
		}
	}

	return fmt.Sprint(sum)
}

func day11Part2(input string) string {

	galaxies := galaxiesAfterExpension(toLines(input), 1_000_000)

	sum := 0
	for i, start := range galaxies[:len(galaxies)-1] {
		for _, end := range galaxies[i+1:] {
			sum += Abs(start.x-end.x) + Abs(start.y-end.y)
		}
	}

	return fmt.Sprint(sum)
}

func Day11(test bool) {
	path := "inputs/2023/day11.txt"
	if test {
		path = strings.Replace(path, "day11", "day11-test", 1)
	}

	input, err := os.ReadFile(path)
	check(err)

	fmt.Println("Day 11")
	fmt.Println("\tPart 1:", day11Part1(string(input)))
	fmt.Println("\tPart 2:", day11Part2(string(input)))
}
