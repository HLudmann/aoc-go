package y2023

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var nodeRe = regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)`)

type Node struct {
	Left, Right string
}

type Network map[string]*Node

func toDirsAndNetwork(input string) (string, Network) {
	dirsAndNetwork := toLines(input)

	dirs := dirsAndNetwork[0]
	network := make(Network, len(dirsAndNetwork)-1)

	for _, node := range dirsAndNetwork[1:] {
		res := nodeRe.FindStringSubmatch(node)
		for _, n := range res[1:] {
			if network[n] == nil {
				network[n] = &Node{}
			}
		}
		node := network[res[1]]
		node.Left = res[2]
		node.Right = res[3]
	}

	return dirs, network
}

func stepsToEnd(network Network, dirs string, start string, isEnd func(pos string) bool) (steps int) {
	pos := start
	for !isEnd(pos) {
		switch dirs[steps%len(dirs)] {
		case 'L':
			pos = network[pos].Left
		case 'R':
			pos = network[pos].Right
		}
		steps++
	}

	return
}

func day08Part1(input string) string {
	dirs, network := toDirsAndNetwork(input)

	return fmt.Sprint(stepsToEnd(network, dirs, "AAA", func(pos string) bool { return pos == "ZZZ" }))
}

func day08Part2(input string) string {
	dirs, network := toDirsAndNetwork(input)

	starts := []string{}
	for key := range network {
		if key[2] == 'A' {
			starts = append(starts, key)
		}
	}

	minSteps := make(chan int, len(starts))

	for _, start := range starts {
		start := start
		go func() {
			minSteps <- stepsToEnd(network, dirs, start, func(pos string) bool { return pos[2] == 'Z' })
		}()
	}

	steps := <-minSteps
	for i := 1; i < len(starts); i++ {
		steps = Lcm(steps, <-minSteps)
	}

	return fmt.Sprint(steps)
}

func Day08(test bool) {
	path := "inputs/2023/day08.txt"
	if test {
		path = strings.Replace(path, "day08", "day08-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day08Part1(string(input))
	p2 := day08Part2(string(input))

	fmt.Printf("Day 08\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
