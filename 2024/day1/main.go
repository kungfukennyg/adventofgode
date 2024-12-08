package day1

import (
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/common"
)

func similarityScore(input string) int {
	left, right := parse(input)
	unique := map[int]struct{}{}
	for _, l := range left {
		unique[l] = struct{}{}
	}

	counts := map[int]int{}
	for _, r := range right {
		if _, ok := unique[r]; !ok {
			continue
		}

		counts[r] += 1
	}

	similarity := 0
	for _, l := range left {
		count := counts[l]
		similarity += l * count
	}
	return similarity
}

func parse(input string) ([]int, []int) {
	lines := strings.Split(string(input), "\n")
	left, right := make([]int, 0, len(lines)), make([]int, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		l := common.MustAtoi(strings.TrimSpace(parts[0]))
		r := common.MustAtoi(strings.TrimSpace(parts[1]))

		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func totalDistance(input string) int {
	left, right := parse(input)
	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i := range left {
		a, b := left[i], right[i]

		total += absDiff(a, b)
	}

	return total
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
