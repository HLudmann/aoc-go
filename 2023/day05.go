package y2023

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type conversionRange struct {
	Destination int
	Source      int
	Size        int
}

func (cr conversionRange) Contains(value int) bool {
	return cr.Source <= value && value < cr.Source+cr.Size
}

func (cr conversionRange) Convert(value int) int {
	return cr.Destination + value - cr.Source
}

func (cr conversionRange) Intersection(sr []int) []int {
	if cr.Source+cr.Size < sr[0] || sr[0]+sr[1] < cr.Source {
		return []int{0, 0}
	}

	start := max(cr.Source, sr[0])

	return []int{
		start,
		min(cr.Source+cr.Size, sr[0]+sr[1]) - start,
	}
}

type conversionMap struct {
	Ranges []conversionRange
}

func (cm conversionMap) Convert(value int) int {
	for _, cr := range cm.Ranges {
		if cr.Contains(value) {
			return cr.Convert(value)
		}
	}
	return value
}

func (cm conversionMap) ConvertRange(seedRange []int) [][]int {
	var ranges [][]int
	start := seedRange[0]
	last := seedRange[0] + seedRange[1] - 1
	for _, cr := range cm.Ranges {
		if inter := cr.Intersection(seedRange); inter[1] != 0 {
			if start < inter[0] {
				ranges = append(ranges, []int{start, inter[0] - start})
			}
			start = inter[0] + inter[1]
			inter[0] = cr.Convert(inter[0])
			ranges = append(ranges, inter)

			if last < start {
				break
			}
		}
	}

	if start < last {
		ranges = append(ranges, []int{start, last - start + 1})
	}
	return ranges
}

func (cm conversionMap) ConvertRanges(seedRanges [][]int) [][]int {
	var ranges [][]int
	for _, sr := range seedRanges {
		ranges = append(ranges, cm.ConvertRange(sr)...)
	}
	return ranges
}

func minimum(locations [][]int) int {
	minLoc := -1
	for _, loc := range locations {
		if minLoc < 0 || loc[0] < minLoc {
			minLoc = loc[0]
		}
	}
	return minLoc
}

func parseSeedAndMaps(input string) (seeds []int, cms []conversionMap) {
	seeds_n_maps := strings.Split(input[:len(input)-1], "\n\n")

	for _, seed_str := range strings.Split(seeds_n_maps[0][7:], " ") {
		seeds = append(seeds, toInt(seed_str))
	}

	for _, map_description := range seeds_n_maps[1:] {
		var cm conversionMap
		for _, range_str := range strings.Split(map_description, "\n")[1:] {
			dest_src_size := strings.Split(range_str, " ")
			cm.Ranges = append(cm.Ranges, conversionRange{
				toInt(dest_src_size[0]),
				toInt(dest_src_size[1]),
				toInt(dest_src_size[2]),
			})
		}
		slices.SortStableFunc(cm.Ranges, func(a, b conversionRange) int { return a.Source - b.Source })
		cms = append(cms, cm)
	}

	return
}

func seedToLocation(maps []conversionMap, seed int) (val int) {
	val = seed
	for _, m := range maps {
		val = m.Convert(val)
	}
	return
}

func seedRangeToLocation(maps []conversionMap, seedRange []int) int {
	locRanges := [][]int{seedRange}

	for _, cm := range maps {
		locRanges = cm.ConvertRanges(locRanges)
	}

	return minimum(locRanges)
}

func day05Part1(input string) string {
	seeds, maps := parseSeedAndMaps(input)

	locations := make(chan int, len(seeds))

	for _, seed := range seeds {
		s := seed
		go func() {
			locations <- seedToLocation(maps, s)
		}()
	}

	min_loc := -1
	for i := 0; i < len(seeds); i++ {
		location := <-locations
		if min_loc == -1 || location < min_loc {
			min_loc = location
		}
	}

	return fmt.Sprint(min_loc)
}

func day05Part2(input string) string {
	seed_info, maps := parseSeedAndMaps(input)

	var seed_ranges [][]int
	for i := 0; i < len(seed_info); i += 2 {
		seed_ranges = append(seed_ranges, []int{seed_info[i], seed_info[i+1]})
	}

	locations := make(chan int, len(seed_ranges))

	for _, seedRange := range seed_ranges {
		sr := seedRange
		go func() {
			locations <- seedRangeToLocation(maps, sr)
		}()
	}

	min_loc := -1
	for i := 0; i < len(seed_ranges); i++ {
		location := <-locations
		if min_loc == -1 || location < min_loc {
			min_loc = location
		}
	}

	return fmt.Sprint(min_loc)
}

func Day05(test bool) {
	path := "inputs/2023/day05.txt"
	if test {
		path = strings.Replace(path, "day05", "day05-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day05Part1(string(input))
	p2 := day05Part2(string(input))

	fmt.Printf("Day 05\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
