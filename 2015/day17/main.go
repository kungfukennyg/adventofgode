package day17

import (
	"math"
	"slices"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func sum(values []int) int {
	s := 0
	for _, n := range values {
		s += n
	}
	return s
}

func uniqueCombos(input string, liters int) (int, int) {
	containers := parse(input)
	slices.Sort(containers)
	slices.Reverse(containers)
	unique := 0
	comboLens := map[int]int{}
	// for n containers, there are 2^n possible combinations
	// i here represents every binary value from 1 -> 2^n
	// with each set bit representing an included container index
	for i := 1; i <= 1<<len(containers); i++ {
		s := 0
		len := 0
		for j, c := range containers {
			if i&(1<<j) > 0 {
				s += c
				len++
			}
		}
		// count only the combinations that exactly equal our
		// desired capacity
		if s == liters {
			unique++
			comboLens[len] += 1
		}
	}
	lowest := math.MaxInt64
	for k := range comboLens {
		lowest = min(lowest, k)
	}
	return unique, comboLens[lowest]
}

func parse(input string) []int {
	out := []int{}
	for _, s := range aoc.Lines(input) {
		out = append(out, aoc.MustAtoi(s))
	}
	return out
}
