package y2023

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Signal struct {
	Pulse    bool
	Src, Dst string
}

type Module struct {
	Type       rune
	State      bool
	Prev, Next []string
	Memory     map[string]bool
}

func (m *Module) AddPrev(p string) {
	if slices.Contains(m.Prev, p) {
		return
	}
	m.Prev = append(m.Prev, p)
}

func (m *Module) AddTypeAndNext(t rune, n []string) {
	m.Type = t
	m.Next = n

	if t == '&' {
		m.Memory = make(map[string]bool)
	} else {
		m.Memory = nil
	}
}

func (m *Module) ProcessSignal(s Signal) (res []Signal) {
	if m.Type == '%' && s.Pulse {
		return
	}

	pulse := s.Pulse
	switch m.Type {
	case '%':
		m.State = !m.State
		pulse = m.State
	case '&':
		m.Memory[s.Src] = pulse
		var low int
		for _, p := range m.Prev {
			if !m.Memory[p] {
				low++
			}
		}
		pulse = low > 0
	}

	for _, n := range m.Next {
		res = append(res, Signal{pulse, s.Dst, n})
	}
	return
}

func NewModule(t rune, p, n []string) *Module {
	m := Module{t, false, p, n, nil}
	if t == '&' {
		m.Memory = make(map[string]bool)
	}
	return &m
}

type ModulesGraph map[string]*Module

func (mg ModulesGraph) ToDotFile(filename string) error {

	g := graph.New(graph.StringHash, graph.Directed())
	for key := range mg {
		g.AddVertex(key)
	}
	for key, val := range mg {
		for _, n := range val.Next {
			g.AddEdge(key, n, graph.EdgeWeight(int(val.Type)))
		}
	}

	file, _ := os.Create(filename)
	return draw.DOT(g, file)
}

func day20Part1(input string) string {
	modules := make(map[string]*Module)
	for _, line := range toLines(input) {
		s := strings.Split(line, " -> ")
		t := rune(s[0][0])
		next := strings.Split(s[1], ", ")
		name := s[0][1:]
		if t == 'b' {
			name = s[0]
		}

		for _, n := range next {
			if m, ok := modules[n]; ok {
				m.AddPrev(name)
			} else {
				modules[n] = NewModule('u', []string{name}, []string{})
			}
		}

		if m, ok := modules[name]; ok {
			m.AddTypeAndNext(t, next)
		} else {
			modules[name] = NewModule(t, []string{}, next)
		}
	}

	var low, high int
	for i := 0; i < 1000; i++ {
		queue := []Signal{{false, "button", "broadcaster"}}
		for len(queue) != 0 {
			s := queue[0]
			if s.Pulse {
				high++
			} else {
				low++
			}
			m := modules[s.Dst]
			queue = append(queue[1:], m.ProcessSignal(s)...)
		}
	}
	fmt.Println("H:", high, "L:", low)
	return fmt.Sprint(high * low)
}

func day20Part2(input string) string {
	modules := make(ModulesGraph)
	for _, line := range toLines(input) {
		s := strings.Split(line, " -> ")
		t := rune(s[0][0])
		next := strings.Split(s[1], ", ")
		name := s[0][1:]
		if t == 'b' {
			name = s[0]
		}

		for _, n := range next {
			if m, ok := modules[n]; ok {
				m.AddPrev(name)
			} else {
				modules[n] = NewModule('u', []string{name}, []string{})
			}
		}

		if m, ok := modules[name]; ok {
			m.AddTypeAndNext(t, next)
		} else {
			modules[name] = NewModule(t, []string{}, next)
		}
	}

	subGraphStarts := modules["broadcaster"].Next
	minimumToRxLow := 1

	for _, start := range subGraphStarts {
		var i int
		sentTrue := false
		for !sentTrue {
			i++
			queue := []Signal{{false, "broadcaster", start}}
			for len(queue) != 0 {
				nq := []Signal{}
				for _, s := range queue {
					if s.Dst == "cn" && s.Pulse {
						sentTrue = true
						break
					}
					nq = append(nq, modules[s.Dst].ProcessSignal(s)...)
				}
				queue = nq
			}
		}
		minimumToRxLow = Lcm(minimumToRxLow, i)
	}

	return fmt.Sprint(minimumToRxLow)
}

func Day20(test bool) {
	path := "inputs/2023/day20.txt"
	if test {
		path = strings.Replace(path, "day20", "day20-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day20Part1(string(input))
	p2 := day20Part2(string(input))

	fmt.Printf("Day 20\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
