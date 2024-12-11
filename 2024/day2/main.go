package day2

import (
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func problemDampener(report []int) bool {
	for i := range report {
		var sub []int
		if i == 0 {
			sub = report[1:]
		} else if i+1 == len(report) {
			sub = report[:len(report)-1]
		} else {
			sub = slices.Concat(report[:i], report[i+1:])
		}

		if isSafe(sub) {
			return true
		}
	}

	return false
}

func countSafe(input string, withDampener bool) int {
	reports := loadReport(input)
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		} else if withDampener && problemDampener(report) {
			safeCount++
		}
	}
	return safeCount
}

func isSafe(report []int) bool {
	asc := false
	for i, level := range report {
		if i == 0 {
			continue
		}
		prev := report[i-1]
		if prev == level {
			return false
		}

		if i == 1 {
			asc = level > prev
		}
		if (prev > level && asc) ||
			(prev < level && !asc) {
			// not solely asc/desc
			return false
		}
		if (prev > level && prev-level > 3) ||
			(prev < level && level-prev > 3) {
			return false
		}
	}
	return true
}

func loadReport(input string) [][]int {
	reports := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		report := []int{}
		for _, level := range strings.Split(line, " ") {
			report = append(report, aoc.MustAtoi(level))
		}
		reports = append(reports, report)
	}
	return reports
}
