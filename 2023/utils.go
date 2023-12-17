package y2023

import (
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

type Pos struct {
	x, y int
}

func (a Pos) Add(b Pos) Pos      { return Pos{a.x + b.x, a.y + b.y} }
func (p Pos) Multiply(n int) Pos { return Pos{p.x * n, p.y * n} }

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

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func TrimLeft(s string, r rune) string {
	if s == "" || rune(s[0]) != r {
		return s
	}

	return TrimLeft(s[1:], r)
}

func Transpose[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 {
		return matrix
	}
	t := make([][]T, len(matrix[0]))
	for _, row := range matrix {
		for j, val := range row {
			t[j] = append(t[j], val)
		}
	}
	return t
}

func TransposeStr(matrix []string) []string {
	if len(matrix) == 0 {
		return matrix
	}
	t := make([]string, len(matrix[0]))
	for _, row := range matrix {
		for j, val := range row {
			t[j] = t[j] + string(val)
		}
	}
	return t
}

func StrDiff(s1, s2 string) (indexes []int) {
	for i := 0; i < min(len(s1), len(s2)); i++ {
		if s1[i] != s2[i] {
			indexes = append(indexes, i)
		}
	}
	return
}
