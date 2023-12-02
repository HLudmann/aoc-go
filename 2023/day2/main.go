package a23d2

import (
	"fmt"
	"strconv"
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

func toInt(str string) int {
	digit, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return digit
}

func parse(input string) (games []Game) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		game_sets := strings.Split(line, ": ")
		sets := strings.Split(game_sets[1], "; ")
		game := Game{toInt(game_sets[0][5:]), make([]Set, len(sets))}
		for i, set_to_split := range sets {
			splited_set := strings.Split(set_to_split, ", ")
			set := make([]Cubes, len(splited_set))
			for j, cubes_to_split := range splited_set {
				amount_color := strings.Split(cubes_to_split, " ")
				set[j] = Cubes{amount_color[1], toInt(amount_color[0])}
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

func Puzzle1(input string) string {
	bag := map[string]int{"red": 12, "green": 13, "blue": 14}
	games := parse(input)
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

func Puzzle2(input string) string {
	games := parse(input)
	var sum int
	for _, game := range games {
		sum += game.Power()
	}
	return fmt.Sprint(sum)
}
