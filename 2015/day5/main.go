package day5

import (
	"strings"
)

const vowels string = "aeiou"

var disallowed []string = []string{"ab", "cd", "pq", "xy"}

func countNice(input string) int {
	niceCount := 0
	for _, line := range strings.Split(input, "\n") {
		if isNice(line) {
			niceCount++
		}
	}
	return niceCount
}

func isNice(line string) bool {
	for _, sub := range disallowed {
		if strings.Contains(line, sub) {
			return false
		}
	}

	prev := rune(line[0])
	twiceInRow := false
	vowelCount := 0
	for i, r := range line {
		if strings.ContainsRune(vowels, r) {
			vowelCount++
		}

		if i == 0 {
			continue
		}
		if prev == r {
			twiceInRow = true
		}
		if twiceInRow && vowelCount >= 3 {
			return true
		}
		prev = r
	}

	return false
}

func countNiceNewModel(input string) int {
	nice := 0
	for _, line := range strings.Split(input, "\n") {
		hasAa, hasAba := false, false
		for i, tok := range line {
			if i+2 >= len(line) {
				break
			}

			next := line[i+1]
			sub := line[i+2:]
			if !hasAa && strings.Contains(sub, string(tok)+string(next)) {
				hasAa = true
			}

			third := line[i+2]
			if !hasAba && rune(third) == tok {
				hasAba = true
			}

			if hasAa && hasAba {
				nice++
				break
			}
		}
	}

	return nice
}
