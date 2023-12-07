package y2023

import (
	"fmt"
	"os"
	"strings"
)

type Cubes struct {
	Colour string
	Amount int
}

type Set []Cubes

type Game struct {
	Id   int
	Sets []Set
}

func parseGames(input string) (games []Game) {
	for _, line := range toLines(input) {
		gameSets := strings.Split(line, ": ")
		sets := strings.Split(gameSets[1], "; ")
		game := Game{toInt(gameSets[0][5:]), make([]Set, len(sets))}
		for i, setToSplit := range sets {
			splitedSet := strings.Split(setToSplit, ", ")
			set := make([]Cubes, len(splitedSet))
			for j, cubesToSplit := range splitedSet {
				amountColor := strings.Fields(cubesToSplit)
				set[j] = Cubes{amountColor[1], toInt(amountColor[0])}
			}
			game.Sets[i] = set
		}
		games = append(games, game)
	}
	return
}

func possibleCubes(bag map[string]int, cubes Cubes) bool {
	return cubes.Amount <= bag[cubes.Colour]
}

func possibleSet(bag map[string]int, set Set) bool {
	for _, cubes := range set {
		if !possibleCubes(bag, cubes) {
			return false
		}
	}
	return true
}

func possibleSets(bag map[string]int, sets []Set) bool {
	for _, set := range sets {
		if !possibleSet(bag, set) {
			return false
		}
	}
	return true
}

func day02Part1(input string) string {
	bag := map[string]int{"red": 12, "green": 13, "blue": 14}
	games := parseGames(input)
	var sum int

	for _, game := range games {
		if possibleSets(bag, game.Sets) {
			sum += game.Id
		}
	}

	return fmt.Sprint(sum)
}

func (g Game) Power() int {
	miniBag := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, set := range g.Sets {
		for _, cubes := range set {
			if miniBag[cubes.Colour] < cubes.Amount {
				miniBag[cubes.Colour] = cubes.Amount
			}
		}
	}
	power := 1
	for _, amount := range miniBag {
		power *= amount
	}
	return power
}

func day02Part2(input string) string {
	games := parseGames(input)
	var sum int
	for _, game := range games {
		sum += game.Power()
	}
	return fmt.Sprint(sum)
}

func Day02(test bool) {
	path := "inputs/2023/day02.txt"
	if test {
		path = strings.Replace(path, "day02", "day02-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day02Part1(string(input))
	p2 := day02Part2(string(input))

	fmt.Printf("Day 02\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
