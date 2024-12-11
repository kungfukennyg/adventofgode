package day5

import (
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type rule struct {
	after map[int]struct{}
}

func correctPageNumbers(input string) (int, int) {
	rules, updates := parse(input)
	correct, wrong := 0, 0

	for _, pages := range updates {
		sorted := sort(rules, pages)

		if slices.Equal(pages, sorted) {
			correct += pages[len(pages)/2]
		} else {
			wrong += sorted[len(pages)/2]
		}
	}
	return correct, wrong
}

func sort(rules map[int]rule, update []int) []int {
	u := slices.Clone(update)
	slices.SortFunc(u, func(a, b int) int {
		if a == b {
			return 0
		}

		if _, ok := rules[a].after[b]; ok {
			return -1
		} else if _, ok := rules[b].after[a]; ok {
			return 1
		}

		return 0
	})

	return u
}

func parse(input string) (map[int]rule, [][]int) {
	rules := map[int]rule{}
	updates := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			before, after := aoc.MustAtoi(parts[0]), aoc.MustAtoi(parts[1])
			r, ok := rules[before]
			if !ok {
				r = rule{
					after: map[int]struct{}{},
				}
				rules[before] = r
			}

			r.after[after] = struct{}{}
		} else if strings.Contains(line, ",") {
			update := []int{}
			for _, page := range strings.Split(line, ",") {
				update = append(update, aoc.MustAtoi(page))
			}
			updates = append(updates, update)
		}
	}
	return rules, updates
}
