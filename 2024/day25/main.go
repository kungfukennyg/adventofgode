package day25

import (
	"github.com/kungfukennyg/adventofgode/aoc"
)

func keyFits(lock, key []int, height int) bool {
	for i, pl := range lock {
		if pl+key[i] > height {
			return false
		}
	}
	return true
}

func lockMatches(input string) int {
	locks, keys, height := parse(input)

	matches := 0
	for _, lock := range locks {
		for _, key := range keys {
			if keyFits(lock, key, height) {
				matches++
			}
		}
	}

	return matches
}

func parse(input string) ([][]int, [][]int, int) {
	sections := [][]string{}

	height := 0
	cur := []string{}
	for _, line := range aoc.Lines(input) {
		if len(line) == 0 {
			sections = append(sections, cur)
			cur = []string{}
			continue
		}

		cur = append(cur, line)
		height = max(height, len(cur))
	}
	if len(cur) > 0 {
		sections = append(sections, cur)
	}

	locks, keys := parseLocksKeys(sections)
	return locks, keys, height - 2
}

func parseLocksKeys(sections [][]string) ([][]int, [][]int) {
	locks, keys := [][]int{}, [][]int{}
	for _, part := range sections {
		var cur []int
		var lock bool
		for y, line := range part {
			if cur == nil {
				lock = aoc.ContainsOnly(line, "#")
				cur = make([]int, len(line))
				continue
			}
			if y+1 >= len(part) {
				continue
			}

			for x, s := range line {
				if s == '#' {
					cur[x]++
				}
			}
		}

		if len(cur) > 0 {
			if lock {
				locks = append(locks, cur)
			} else {
				keys = append(keys, cur)
			}
			cur = nil
		}
	}
	return locks, keys
}
