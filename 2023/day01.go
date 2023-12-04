package y2023

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day01Part1(input string) string {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`[a-zA-Z]`)

	var sum int
	for _, line := range lines {
		if line == "" {
			continue
		}
		only_digits := re.ReplaceAllString(line, "")
		first_digit := toDigit(string(only_digits[0]))
		last_digit := toDigit(string(only_digits[len(only_digits)-1]))

		sum += first_digit*10 + last_digit
	}

	return fmt.Sprint(sum)
}

func toDigit(str string) int {
	digit, err := strconv.Atoi(str)
	check(err)
	return digit
}

func day01Part2(input string) string {
	for key, val := range map[string]string{"one": "o1e", "two": "t2o", "three": "t3e", "four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e"} {
		re := regexp.MustCompile(fmt.Sprint(key))
		input = re.ReplaceAllString(input, val)
	}
	return day01Part1(input)
}

func Day01(test bool) {
	path := "inputs/2023/day01.txt"
	if test {
		path = strings.Replace(path, "day01", "day01-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	d2p1 := day01Part1(string(input))
	d2p2 := day01Part2(string(input))

	fmt.Printf("Day 01\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", d2p1, d2p2)
}