package y2023

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Pos struct {
	x, y int
}

func findAdjacentNumbers(lines []string, used map[Pos]bool, p Pos) (numbers []int) {
	xRange := []int{p.x}
	if p.x > 0 {
		xRange = append(xRange, p.x-1)
	}
	if p.x < len(lines)-1 {
		xRange = append(xRange, p.x+1)
	}
	re := regexp.MustCompile(`\d+`)
	for _, x := range xRange {
		matches := re.FindAllStringIndex(lines[x], -1)
		for _, match := range matches {
			if used[Pos{x, match[0]}] {
				continue
			}
			if p.y < match[0]-1 || match[1] < p.y {
				continue
			}
			number := toInt(lines[x][match[0]:match[1]])
			for y := match[0]; y < match[1]; y++ {
				used[Pos{x, y}] = true
			}
			numbers = append(numbers, number)
		}
	}
	return
}

func day03Part1(input string) string {
	lines := toLines(input)
	var sum int
	used := make(map[Pos]bool)
	symbolRe := regexp.MustCompile(`[^\d\n.]`)
	for x, line := range lines {
		symbolYs := symbolRe.FindAllStringIndex(line, -1)
		for _, y := range symbolYs {
			numbers := findAdjacentNumbers(lines, used, Pos{x, y[0]})
			for _, number := range numbers {
				sum += number
			}
		}
	}
	return fmt.Sprint(sum)
}

func day03Part2(input string) string {
	lines := toLines(input)
	var sum int
	used := make(map[Pos]bool)
	gearRe := regexp.MustCompile(`\*`)
	for x, line := range lines {
		symbolYs := gearRe.FindAllStringIndex(line, -1)
		for _, y := range symbolYs {
			numbers := findAdjacentNumbers(lines, used, Pos{x, y[0]})
			if len(numbers) == 2 {
				sum += numbers[0] * numbers[1]
			}
		}
	}
	return fmt.Sprint(sum)
}

func Day03(test bool) {
	path := "inputs/2023/day03.txt"
	if test {
		path = strings.Replace(path, "day03", "day03-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day03Part1(string(input))
	p2 := day03Part2(string(input))

	fmt.Printf("Day 03\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
