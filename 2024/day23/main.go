package day23

import (
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func setToPass(path aoc.Set[*aoc.Vertex[string]]) string {
	pass := make([]string, 0, len(path))
	for v := range path.Values() {
		pass = append(pass, v.Key)
	}
	slices.Sort(pass)
	return strings.Join(pass, ",")
}

func lanPassword(input string) string {
	graph := parse(input)

	cliques := graph.Cliques()
	var longest aoc.Set[*aoc.Vertex[string]]
	for _, cliq := range cliques {
		if longest == nil || len(cliq) > len(longest) {
			longest = cliq
		}
	}

	return setToPass(longest)
}

func filterMatches(set aoc.Set[*aoc.Vertex[string]], filter string) bool {
	if filter == "" {
		return true
	}

	for v := range set {
		if strings.HasPrefix(v.Key, filter) {
			return true
		}
	}
	return false
}

func interconnected(input, filter string, depth int) int {
	graph := parse(input)
	cliques := graph.CliquesN(depth)
	sum := 0
	for _, clique := range cliques {
		if filterMatches(clique, filter) {
			sum++
		}
	}
	return sum
}

func parse(input string) *aoc.Graph[string] {
	graph := aoc.NewGraph[string]()

	for _, line := range aoc.Lines(input) {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		graph.AddVertex(a, a)
		graph.AddVertex(b, b)
		graph.AddEdge(a, b, 1)
		graph.AddEdge(b, a, 1)
	}
	return graph
}
