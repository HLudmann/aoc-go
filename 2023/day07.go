package y2023

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Hand struct {
	Cards     []int
	Bid, Type int
}

func NewHand(cards []int, bid int) Hand {
	counts := make(map[int]int)
	for _, card := range cards {
		counts[card] += 1
	}
	t := 0.0
	for _, val := range counts {
		t += math.Pow(10, float64(val-1))
	}

	return Hand{cards, bid, int(t)}
}

func NewHandWithJoker(cards []int, bid int) Hand {
	counts := make(map[int]int)
	js := 0
	for _, card := range cards {
		if card == 1 {
			js++
			continue
		}
		counts[card] += 1
	}
	t := 0.0
	for _, val := range counts {
		t += math.Pow(10, float64(val-1))
	}

	if js > 0 {
		switch {
		case 1000 <= t:
			t = 10000
		case 100 <= t:
			t = t - 100 + math.Pow(10, float64(js+2))
		case 10 <= t:
			t = t - 10 + math.Pow(10, float64(js+1))
		case 1 <= t:
			t = t - 1 + math.Pow(10, float64(js))
		case t == 0:
			t = 10000
		}
	}

	return Hand{cards, bid, int(t)}
}

func handsDiff(h1, h2 Hand) int {
	if td := h1.Type - h2.Type; td != 0 {
		return td
	}
	for i := 0; i < 5; i++ {
		if d := h1.Cards[i] - h2.Cards[i]; d != 0 {
			return d
		}
	}
	return 0
}

func cardToInt(card string, joker bool) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if joker {
			return 1
		}
		return 11
	case "T":
		return 10
	default:
		return toInt(card)
	}
}

func day07Part1(input string) string {
	var hands []Hand
	for _, line := range toLines(input) {
		cardsAndBid := strings.Fields(line)
		bid := toInt(cardsAndBid[1])

		cards := make([]int, 5)
		for i, card := range cardsAndBid[0] {
			cards[i] = cardToInt(string(card), false)
		}

		hands = append(hands, NewHand(cards, bid))
	}

	slices.SortStableFunc(hands, handsDiff)

	var sum int
	for i, hand := range hands {
		sum += (i + 1) * hand.Bid
	}

	return fmt.Sprint(sum)
}

func day07Part2(input string) string {
	var hands []Hand
	for _, line := range toLines(input) {
		cardsAndBid := strings.Fields(line)
		bid := toInt(cardsAndBid[1])

		cards := make([]int, 5)
		for i, card := range cardsAndBid[0] {
			cards[i] = cardToInt(string(card), true)
		}

		hands = append(hands, NewHandWithJoker(cards, bid))
	}

	slices.SortStableFunc(hands, handsDiff)

	var sum int
	for i, hand := range hands {
		sum += (i + 1) * hand.Bid
	}

	return fmt.Sprint(sum)
}

func Day07(test bool) {
	path := "inputs/2023/day07.txt"
	if test {
		path = strings.Replace(path, "day07", "day07-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day07Part1(string(input))
	p2 := day07Part2(string(input))

	fmt.Printf("Day 07\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
