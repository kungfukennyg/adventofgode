package day9

import (
	"log"
	"math"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type graph struct {
	vertices map[string]*vertex
}

type edge struct {
	weight int
	v      *vertex
}

type vertex struct {
	country string
	edges   map[string]*edge
}

func (g *graph) addVertex(src string) {
	if _, ok := g.vertices[src]; ok {
		return
	}

	g.vertices[src] = &vertex{country: src, edges: map[string]*edge{}}
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

func (g *graph) route(r []string) (int, bool) {
	if len(r) < 2 && len(r) < 1 {
		return 0, false
	}

	dist := 0
	src := r[0]
	for _, to := range r[1:] {
		from, ok := g.vertices[src]
		if !ok {
			return 0, false
		}
		e, ok := from.edges[to]
		if !ok {
			return 0, false
		}
		dist += e.weight
		src = to
	}
	return dist, true
}

func (g *graph) keys() []string {
	o := make([]string, 0, len(g.vertices))
	for k := range g.vertices {
		o = append(o, k)
	}
	return o
}

func longestDist(input string) int {
	g := parse(input)
	countries := g.keys()
	perms := aoc.Permutate(countries)

	longest := math.MinInt64
	for _, route := range perms {
		dist, ok := g.route(route)
		if ok && dist > 0 {
			longest = max(longest, dist)
			// log.Printf("%s = %d\n", strings.Join(route, " -> "), dist)
		}
	}
	return longest
}

func shortestDist(input string) int {
	g := parse(input)
	countries := g.keys()
	perms := aoc.Permutate(countries)

	lowest := math.MaxInt64
	for _, route := range perms {
		dist, ok := g.route(route)
		if ok && dist > 0 {
			lowest = min(lowest, dist)
			log.Printf("%s = %d\n", strings.Join(route, " -> "), dist)
		}
	}
	return lowest
}

func parse(input string) *graph {
	g := graph{
		vertices: map[string]*vertex{},
	}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		from, to := parts[0], parts[2]
		dist := aoc.MustAtoi(parts[4])

		g.addVertex(from)
		g.addVertex(to)

		g.addEdge(from, to, dist)
		g.addEdge(to, from, dist)
	}

	return &g
}
