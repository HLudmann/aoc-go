package y2023

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type card struct {
	Id             int
	winningNumbers map[int]bool
	numbers        []int
}

func (c card) Winners() (w int) {
	for _, number := range c.numbers {
		if c.winningNumbers[number] {
			w++
		}
	}
	return
}

func (c card) Worth() int {
	return int(math.Pow(2, float64(c.Winners()-1)))
}

func parseCard(card_str string) card {
	card_win_num := strings.Split(card_str, ": ")
	card_id := toInt(strings.Fields(card_win_num[0])[1])

	winNum_numbers := strings.Split(card_win_num[1], " | ")
	winningNumbers := make(map[int]bool)
	for _, winNum := range strings.Fields(winNum_numbers[0]) {
		winningNumbers[toInt(winNum)] = true
	}
	var numbers []int
	for _, number := range strings.Fields(winNum_numbers[1]) {
		numbers = append(numbers, toInt(number))
	}

	return card{card_id, winningNumbers, numbers}
}

func parseCards(input string) (cards []card) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		cards = append(cards, parseCard(line))
	}
	return
}

func day04Part1(input string) string {
	var sum int
	for _, c := range parseCards(input) {
		sum += c.Worth()
	}
	return fmt.Sprint(sum)
}

func day04Part2(input string) string {
	var sum int
	cards := parseCards(input)
	buffer := make([]int, len(cards[0].winningNumbers))
	for i := 0; i < len(cards[0].winningNumbers); i++ {
		buffer[i] = 1
	}
	for _, c := range cards {
		nbr_card := buffer[0]
		buffer = append(buffer[1:], 1)
		for i := 0; i < c.Winners(); i++ {
			buffer[i] += nbr_card
		}
		sum += nbr_card
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
	p1 := day04Part1(string(input))
	p2 := day04Part2(string(input))

	fmt.Printf("Day 04\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
