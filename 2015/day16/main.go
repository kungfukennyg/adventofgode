package day16

import (
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func whichSue(input, compounds string, greater, lesser []string) int {
	people := parsePeople(input)
	comps := parseCompounds(compounds)

outer:
	for i, p := range people {
		for comp, want := range comps {
			got, ok := p[comp]
			if !ok {
				continue
			}
			if ok := slices.Contains(greater, comp); ok {
				if got <= want {
					continue outer
				}
			} else if ok := slices.Contains(lesser, comp); ok {
				if got >= want {
					continue outer
				}
			} else if want != got {
				continue outer
			}
		}

		return i + 1
	}
	return -1
}

func parsePeople(input string) []map[string]int {
	people := []map[string]int{}
	for _, line := range aoc.Lines(input) {
		person := parsePerson(line)
		people = append(people, person)
	}
	return people
}

func parsePerson(line string) map[string]int {
	out := map[string]int{}
	skip := strings.Index(line, ": ") + 1
	line = line[skip:]
	for _, p := range strings.Split(line, ",") {
		parts := strings.Split(p, ":")
		prop := strings.TrimSpace(parts[0])
		out[prop] = aoc.MustAtoi(strings.TrimSpace(parts[1]))
	}
	return out
}

func parseCompounds(compounds string) map[string]int {
	o := map[string]int{}
	for _, line := range aoc.Lines(compounds) {
		parts := strings.Split(line, ":")
		o[parts[0]] = aoc.MustAtoi(strings.TrimSpace(parts[1]))
	}
	return o
}
