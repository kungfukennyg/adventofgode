package day2

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type keypad struct {
	g   *aoc.Grid[string]
	cur aoc.Pos
}

func (k *keypad) step(move aoc.Dir) {
	next := k.cur.Step(move)
	if k.g.BoundsCheck(next) && k.g.Get(next) != "" {
		k.cur = next
	}
}

func findKeys(kpStr string, start, input string) string {
	moves := parseMoves(input)
	kp := parseKeypad(kpStr, start)

	codes := []string{}
	for _, row := range moves {
		for _, st := range row {
			kp.step(st)
		}

		code := kp.g.Get(kp.cur)
		codes = append(codes, code)
	}

	return strings.Join(codes, "")
}

func parseKeypad(str, start string) *keypad {
	lines := aoc.Lines(str)
	y := len(lines)
	var x int
	for _, l := range lines {
		x = max(x, len(l))
	}

	g := aoc.NewGrid[string](y, x)

	var startPos aoc.Pos
	for y, line := range lines {
		for x, part := range line {
			if part == ' ' {
				continue
			}

			g.SetByXY(x, y, string(part))
			if string(part) == start {
				startPos.X = x
				startPos.Y = y
			}
		}
	}

	return &keypad{
		g:   g,
		cur: startPos,
	}
}

func parseMoves(input string) [][]aoc.Dir {
	allMoves := [][]aoc.Dir{}
	for _, line := range aoc.Lines(input) {
		moves := []aoc.Dir{}
		for _, b := range line {
			var m aoc.Dir
			switch b {
			case 'U':
				m = aoc.DirUp
			case 'L':
				m = aoc.DirLeft
			case 'R':
				m = aoc.DirRight
			case 'D':
				m = aoc.DirDown
			}
			moves = append(moves, m)
		}
		allMoves = append(allMoves, moves)
	}

	return allMoves
}
