package day2

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/common"
)

func ribbonLength(input string) int {
	ribbon := 0
	for _, line := range strings.Split(input, "\n") {
		l, w, h := parseBox(line)
		x, y := smallestSides(l, w, h)

		// perimeter of smallest face + cubic volume
		ribbon += (2*x + 2*y) +
			(l * w * h)
	}
	return ribbon
}

func wrappingPaper(input string) int {
	paper := 0
	for _, line := range strings.Split(input, "\n") {
		l, w, h := parseBox(line)
		x, y := smallestSides(l, w, h)
		area := x * y

		// surface area + area of smallest side
		paper += (2 * l * w) +
			(2 * w * h) +
			(2 * h * l) +
			area
	}
	return paper
}

func parseBox(line string) (int, int, int) {
	parts := strings.Split(line, "x")
	return common.MustAtoi(parts[0]),
		common.MustAtoi(parts[1]),
		common.MustAtoi(parts[2])
}

func smallestSides(l, w, h int) (int, int) {
	c := max(l, w, h)
	if c == l {
		return w, h
	} else if c == w {
		return l, h
	} else {
		return l, w
	}
}
