package y2023

import (
	"fmt"
	"os"
	"strings"
)

func winingHands(time, distance int) int {
	fChan := make(chan int)
	lChan := make(chan int)

	go func() {
		first := 1
		for first < time {
			if (time-first)*first > distance {
				break
			}
			first++
		}
		fChan <- first
	}()

	go func() {
		last := time
		for 0 < last {
			if (time-last)*last > distance {
				break
			}
			last--
		}
		lChan <- last
	}()

	return <-lChan - <-fChan + 1
}

func day06Part1(input string) string {
	timesAndDistances := strings.Split(input[:len(input)-1], "\n")
	times := strings.Fields(timesAndDistances[0])[1:]
	distances := strings.Fields(timesAndDistances[1])[1:]

	prod := 1

	for i := 0; i < len(times); i++ {
		prod *= winingHands(toInt(times[i]), toInt(distances[i]))
	}

	return fmt.Sprint(prod)
}

func day06Part2(input string) string {
	timeAndDistance := strings.Split(input[:len(input)-1], "\n")

	time := toInt(strings.Replace(timeAndDistance[0][5:], " ", "", -1))
	distance := toInt(strings.Replace(timeAndDistance[1][9:], " ", "", -1))

	return fmt.Sprint(winingHands(time, distance))
}

func Day06(test bool) {
	path := "inputs/2023/day06.txt"
	if test {
		path = strings.Replace(path, "day06", "day06-test", 1)
	}

	input, err := os.ReadFile(path)
	check(err)

	fmt.Println("Day 06")
	fmt.Println("\tPart 1:", day06Part1(string(input)))
	fmt.Println("\tPart 2:", day06Part2(string(input)))
}
