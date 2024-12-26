package day13

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type machine struct {
	a, b  aoc.Pos
	prize aoc.Pos
}

func minTokens(input string, prizeMod int) int {
	machines := parse(input, prizeMod)
	costA, costB := 3, 1
	total := 0
	for _, m := range machines {
		ax, ay := m.a.X, m.a.Y
		bx, by := m.b.X, m.b.Y
		px, py := m.prize.X, m.prize.Y

		b := (ay*px - ax*py) / (ay*bx - ax*by)
		br := (ay*px - ax*py) % (ay*bx - ax*by)
		a := (px - b*bx) / ax
		ar := (px - b*bx) % ax

		if ar == 0 && br == 0 {
			total += a*costA + (b * costB)
		}
	}
	return total
}

func parse(input string, prizeMod int) []machine {
	machines := []machine{}
	var cur *machine
	for _, line := range aoc.Lines(input) {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Split(line, ":")
		lhs, rhs := parts[0], parts[1]
		if strings.Contains(lhs, "Button") {
			parts = strings.Split(strings.TrimSpace(rhs), ", ")
			x := aoc.MustAtoi(strings.ReplaceAll(parts[0], "X+", ""))
			y := aoc.MustAtoi(strings.ReplaceAll(parts[1], "Y+", ""))
			if strings.Contains(lhs, "A") {
				cur = &machine{
					a: aoc.Pos{X: x, Y: y},
				}
			} else {
				cur.b = aoc.Pos{X: x, Y: y}
			}
		} else {
			parts = strings.Split(strings.TrimSpace(rhs), ", ")
			x := aoc.MustAtoi(strings.ReplaceAll(parts[0], "X=", ""))
			y := aoc.MustAtoi(strings.ReplaceAll(parts[1], "Y=", ""))
			cur.prize = aoc.Pos{X: x + prizeMod, Y: y + prizeMod}
			machines = append(machines, *cur)
		}
	}
	return machines
}
