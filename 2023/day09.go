package y2023

import (
	"fmt"
	"os"
	"strings"
)

func lineToValues(line string) (values []int) {
	for _, valStr := range strings.Fields(line) {
		values = append(values, toInt(valStr))
	}
	return
}

func toStepsAndAbsSum(values []int) ([]int, int) {
	sum := 0
	steps := make([]int, len(values)-1)
	for i := 1; i < len(values); i++ {
		step := values[i] - values[i-1]
		sum += Abs(step)
		steps[i-1] = step
	}
	return steps, sum
}

func extrapolateNext(values []int) int {
	if len(values) < 1 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	steps, sum := toStepsAndAbsSum(values)
	lastVal := values[len(values)-1]
	if sum == 0 {
		return lastVal
	}
	return lastVal + extrapolateNext(steps)
}

func extrapolatePrevious(values []int) int {
	if len(values) < 1 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	steps, sum := toStepsAndAbsSum(values)
	if sum == 0 {
		return values[0]
	}
	return values[0] - extrapolatePrevious(steps)
}

func day09Part1(input string) string {
	lines := toLines(input)
	predictions := make(chan int, len(lines))

	for _, line := range lines {
		l := line
		go func() {
			values := lineToValues(l)
			predictions <- extrapolateNext(values)
		}()
	}

	var sum int
	for i := 0; i < len(lines); i++ {
		sum += <-predictions
	}
	return fmt.Sprint(sum)
}

func day09Part2(input string) string {
	lines := toLines(input)
	predictions := make(chan int, len(lines))

	for _, line := range lines {
		l := line
		go func() {
			values := lineToValues(l)
			predictions <- extrapolatePrevious(values)
		}()
	}

	var sum int
	for i := 0; i < len(lines); i++ {
		sum += <-predictions
	}
	return fmt.Sprint(sum)
}

func Day09(test bool) {
	path := "inputs/2023/day09.txt"
	if test {
		path = strings.Replace(path, "day09", "day09-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day09Part1(string(input))
	p2 := day09Part2(string(input))

	fmt.Printf("Day 09\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
