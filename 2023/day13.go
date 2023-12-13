package y2023

import (
	"fmt"
	"os"
	"strings"
)

func validateReflectionWithSmudges(pattern []string, index, smudges int) bool {
	for j := 1; j <= min(index, len(pattern)-index-2); j++ {
		diff := StrDiff(pattern[index-j], pattern[index+j+1])
		if len(diff) > smudges {
			return false
		}
		if len(diff) > 0 {
			smudges -= len(diff)
		}
	}
	return smudges == 0
}

func findReflectionLine(pattern []string, smudges int) int {
	for i := 0; i < len(pattern)-1; i++ {
		s := smudges
		diff := StrDiff(pattern[i], pattern[i+1])
		if len(diff) > s {
			continue
		}
		if len(diff) > 0 {
			s -= len(diff)
		}
		if validateReflectionWithSmudges(pattern, i, s) {
			return i + 1
		}
	}
	return 0
}

func findReflection(pattern []string, smudges int) int {

	if line := findReflectionLine(pattern, smudges); line > 0 {
		return line * 100
	}
	return findReflectionLine(TransposeStr(pattern), smudges)
}

func day13Part1(input string) string {
	var sum int
	for _, p := range strings.Split(input, "\n\n") {
		sum += findReflection(toLines(p), 0)
	}

	return fmt.Sprint(sum)
}

func day13Part2(input string) string {
	var sum int
	for _, pattern := range strings.Split(input, "\n\n") {
		sum += findReflection(toLines(pattern), 1)
	}

	return fmt.Sprint(sum)
}

func Day13(test bool) {
	path := "inputs/2023/day13.txt"
	if test {
		path = strings.Replace(path, "day13", "day13-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day13Part1(string(input))
	p2 := day13Part2(string(input))

	fmt.Printf("Day 13\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
