package day13

import (
	"math"
	"strings"

	"github.com/kungfukennyg/adventofgode/common"
)

type graph struct {
	vertices map[string]*vertex
	keys     []string
}	

type edge struct {
	weight int
	v      *vertex
}

type vertex struct {
	person string
	edges  map[string]*edge
}

func (g *graph) addVertex(src string) {
	if _, ok := g.vertices[src]; ok {
		return
	}

	g.vertices[src] = &vertex{person: src, edges: map[string]*edge{}}
	g.keys = append(g.keys, src)
}

func (g *graph) addEdge(src, dst string, weight int) {
	from, ok := g.vertices[src]
	if !ok {
		return
	}

	to, ok := g.vertices[dst]
	if !ok {
		return
	}

	from.edges[dst] = &edge{weight: weight, v: to}
}

func (g *graph) sum(keys []string) int {
	if len(keys) < 2 {
		return 0
	}

	total := 0
	// alice -> bob -> carol -> david
	v := g.vertices[keys[0]]
	for i := 0; i < len(keys); i++ {
		var k string
		if i+1 < len(keys) {
			k = keys[i+1]
		} else {
			k = keys[0]
		}

		e := v.edges[k]
		// a -> b
		total += e.weight
		// b <- a
		total += e.v.edges[v.person].weight
		v = e.v
	}
	return total
}

type opt func(*graph)

func withPerson(person string, happiness int) func(*graph) {
	return func(g *graph) {
		g.addVertex(person)
		for _, p := range g.keys {
			g.addEdge(p, person, happiness)
			g.addEdge(person, p, happiness)
		}
	}
}

func bestSeats(input string, opts ...opt) int {
	g := parse(input)
	for _, o := range opts {
		o(g)
	}

	perms := common.Permutate(g.keys)
	highest := math.MinInt64
	for _, seats := range perms {
		happiness := g.sum(seats)
		if happiness != 0 {
			highest = max(highest, happiness)
		}
	}
	return highest
}

func parse(input string) *graph {
	g := &graph{
		vertices: map[string]*vertex{},
		keys:     []string{},
	}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		a, b := parts[0], parts[10]
		b = b[:len(b)-1]
		happiness := common.MustAtoi(parts[3])
		if parts[2] == "lose" {
			happiness *= -1
		}

		g.addVertex(a)
		g.addVertex(b)
		g.addEdge(a, b, happiness)
	}

	return g
}
