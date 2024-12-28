package day19

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func matchPatterns(input string) (int, int) {
	patterns, desired := parse(input)

	var cache = map[string]int{}
	var matches func(string) int
	matches = func(design string) (n int) {
		if n, ok := cache[design]; ok {
			return n
		}
		defer func() { cache[design] = n }()

		if design == "" {
			return 1
		}

		for p := range patterns {
			if strings.HasPrefix(design, p) {
				n += matches(design[len(p):])
			}
		}

		return n
	}

	possible, ways := 0, 0
	for _, design := range desired {
		n := matches(design)
		if n > 0 {
			possible++
			ways += n
		}
	}

	return possible, ways
}

func parse(input string) (aoc.Set[string], []string) {
	patterns, desired := aoc.Set[string]{}, []string{}
	for _, line := range aoc.Lines(input) {
		if strings.Contains(line, ",") {
			for _, p := range strings.Split(line, ", ") {
				patterns.Add(p)
			}
			continue
		}

		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		desired = append(desired, line)
	}
	return patterns, desired
}
