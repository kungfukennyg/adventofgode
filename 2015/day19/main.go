package day19

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func minSteps(input, startMol string) int {
	dest, graph := parseV2(input)

	steps := 0
	s := ""
	for _, r := range dest {
		if len(s) == 0 {
			s = string(r)
		} else {
			s += string(r)
		}

		// for each piece, try to find a path in graph, if
		// not present, step up the current search str size
		path := graph.DFS(startMol, s)
		if len(path) == 0 {
			continue
		}

		// found path
		steps += len(path) - 1
		s = ""
	}
	return steps
}

func parseV2(input string) (string, *aoc.Graph[string]) {
	var molecule string
	graph := &aoc.Graph[string]{
		Vertices: map[string]*aoc.Vertex[string]{},
		Keys:     []string{},
	}
	for _, line := range aoc.Lines(input) {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if !strings.Contains(line, "=>") {
			molecule = line
			continue
		}

		parts := strings.Split(line, " => ")
		mol, replace := parts[0], parts[1]
		graph.AddVertex(mol, mol)
		graph.AddVertex(replace, replace)
		graph.AddEdge(mol, replace, 1)
	}

	return molecule, graph
}

func distinctMolecules(input string) int {
	start, swaps := parse(input)
	created := aoc.Set[string]{}
	for m, replacements := range swaps {
		indices := aoc.IndicesOf(start, m)
		for _, swap := range replacements {
			for _, i := range indices {
				// HOH
				// O12

				// HHHH
				// 0123
				str := start[:i] + swap + start[i+len(m):]
				created.Add(str)
			}
		}
	}
	return len(created)
}

func parse(input string) (string, map[string][]string) {
	var molecule string
	swaps := map[string][]string{}
	for _, line := range aoc.Lines(input) {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if !strings.Contains(line, "=>") {
			molecule = line
			continue
		}

		parts := strings.Split(line, " => ")
		mol, replace := parts[0], parts[1]
		s, ok := swaps[mol]
		if !ok {
			s = []string{}
		}
		s = append(s, replace)
		swaps[mol] = s
	}
	return molecule, swaps
}
