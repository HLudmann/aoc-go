package y2023

import (
	"strconv"
	"strings"
)

func toDigit(str string) int {
	digit, err := strconv.Atoi(str)
	check(err)
	return digit
}

func toInt(str string) int {
	digit, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return digit
}

func toLines(input string) (lines []string) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return
}

func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}
