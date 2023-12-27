package y2023

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/dominikbraun/graph"
)

func toComponentGraph(input string) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)

	putIn := func(g map[string]map[string]bool, k, v string) {
		if g[k] == nil {
			g[k] = make(map[string]bool)
		}
		g[k][v] = true
	}

	for _, line := range toLines(input) {
		f := strings.Fields(line)
		key := f[0][:len(f[0])-1]
		for _, val := range f[1:] {
			putIn(graph, key, val)
			putIn(graph, val, key)
		}
	}

	return graph
}

func day25Part1(input string) string {
	cg := toComponentGraph(input)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	var vertexes []string
	g := graph.New(graph.StringHash)
	for key := range cg {
		g.AddVertex(key)
		if r.Intn(24) < 7 {
			vertexes = append(vertexes, key)
		}

	}
	for key, val := range cg {
		for n := range val {
			g.AddEdge(key, n)
		}
	}

	getKey := func(a, b string) string {
		if a < b {
			return a + "-" + b
		}
		return b + "-" + a
	}

	freqs := make(map[string]int)
	for i := 0; i < len(vertexes)/200+1; i++ {
		s := min(200, len(vertexes)-200*i)
		c := make(chan map[string]int, s)

		for j, n := range vertexes[i*200 : i*200+s] {
			j := j
			n1 := n
			go func() {
				i := i
				f := make(map[string]int)
				for _, n2 := range vertexes[i*200+j+1:] {
					if n1 == n2 || cg[n1][n2] {
						continue
					}
					path, _ := graph.ShortestPath(g, n1, n2)
					if len(path) < 3 {
						continue
					}
					for i, m1 := range path[:len(path)-1] {
						for _, m2 := range path[i+1:] {
							key := getKey(m1, m2)
							f[key] += 1
						}
					}
				}
				c <- f
			}()
		}

		for j := 0; j < s; j++ {
			for k, v := range <-c {
				freqs[k] += v
			}
		}
	}

	fList := make([][2]interface{}, 0)
	for k, v := range freqs {
		fList = append(fList, [2]interface{}{v, k})
	}
	slices.SortStableFunc(fList, func(a, b [2]interface{}) int { return b[0].(int) - a[0].(int) })

	e1, e2, e3 := fList[0][1].(string), fList[1][1].(string), fList[2][1].(string)

	g2 := graph.New(graph.StringHash)
	for key := range cg {
		g2.AddVertex(key)
	}
	for key, val := range cg {
		for n := range val {
			k := getKey(key, n)
			if k == e1 || k == e2 || k == e3 {
				continue
			}
			g2.AddEdge(key, n)
		}
	}

	ref := e1[:3]
	var l1, l2 int
	for k := range cg {
		_, err := graph.ShortestPath(g2, ref, k)
		if err != nil {
			l1++
		} else {
			l2++
		}
	}

	return fmt.Sprint(l1 * l2)
}

func day25Part2(input string) string {
	return "Push The Big Red Button"
}

func Day25(test bool) {
	path := "inputs/2023/day25.txt"
	if test {
		path = strings.Replace(path, "day25", "day25-test", 1)
	}

	input, err := os.ReadFile(path)
	check(err)

	fmt.Println("Day 25")
	fmt.Println("\tPart 1:", day25Part1(string(input)))
	fmt.Println("\tPart 2:", day25Part2(string(input)))
}
