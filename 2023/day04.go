package y2023

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Card struct {
	Id             int
	winningNumbers map[int]bool
	numbers        []int
}

func (c Card) Winners() (w int) {
	for _, number := range c.numbers {
		if c.winningNumbers[number] {
			w++
		}
	}
	return
}

func (c Card) Worth() int {
	return int(math.Pow(2, float64(c.Winners()-1)))
}

func parseCard(card_str string) Card {
	cardIdAndNum := strings.Split(card_str, ": ")
	cardId := toInt(strings.Fields(cardIdAndNum[0])[1])

	winNumAndNumbers := strings.Split(cardIdAndNum[1], " | ")
	winningNumbers := make(map[int]bool)
	for _, winNum := range strings.Fields(winNumAndNumbers[0]) {
		winningNumbers[toInt(winNum)] = true
	}
	var numbers []int
	for _, number := range strings.Fields(winNumAndNumbers[1]) {
		numbers = append(numbers, toInt(number))
	}

	return Card{cardId, winningNumbers, numbers}
}

func day04Part1(input string) string {
	var sum int
	for _, line := range toLines(input) {
		sum += parseCard(line).Worth()
	}
	return fmt.Sprint(sum)
}

func day04Part2(input string) string {
	var sum int
	buffer := make([]int, 10)
	for i := 0; i < 10; i++ {
		buffer[i] = 1
	}
	for _, line := range toLines(input) {
		nbrCard := buffer[0]
		buffer = append(buffer[1:], 1)
		for i := 0; i < parseCard(line).Winners(); i++ {
			buffer[i] += nbrCard
		}
		sum += nbrCard
	}

	return fmt.Sprint(sum)
}

func Day04(test bool) {
	path := "inputs/2023/day04.txt"
	if test {
		path = strings.Replace(path, "day04", "day04-test", 1)
	}

	input, err := os.ReadFile(path)
	check(err)

	fmt.Println("Day 04")
	fmt.Println("\tPart 1:", day04Part1(string(input)))
	fmt.Println("\tPart 2:", day04Part2(string(input)))
}
