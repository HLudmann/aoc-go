package a23d1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func Puzzle1(input string) string {
	lines := parse(input)

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

func Puzzle2(input string) string {
	for key, val := range map[string]string{"one": "o1e", "two": "t2o", "three": "t3e", "four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e"} {
		re := regexp.MustCompile(fmt.Sprint(key))
		input = re.ReplaceAllString(input, val)
	}
	return Puzzle1(input)
}
