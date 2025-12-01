package day1

import (
	"github.com/kungfukennyg/adventofgode/aoc"
)

func simulate(input string) int {
	n, start := 0, 50
	for _, line := range aoc.Lines(input) {
		step := aoc.MustAtoi(line[1:])
		switch line[0] {
		case 'L':
			start -= step
		case 'R':
			// 95 + 15
			start += step
		}

		if start%100 == 0 {
			n++
		}
	}

	return n
}

func simulatePartTwo(input string) int {
	n, start := 0, 50
	for _, line := range aoc.Lines(input) {
		step := aoc.MustAtoi(line[1:])
		switch line[0] {
		case 'L':
			for range step {
				start -= 1
				if start%100 == 0 {
					n++
				}
			}
		case 'R':
			for range step {
				start += 1
				if start%100 == 0 {
					n++
				}
			}
		}
	}

	return n
}
