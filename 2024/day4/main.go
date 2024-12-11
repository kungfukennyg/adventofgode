package day4

import (
	"fmt"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type dir int

const (
	dirLeft dir = iota
	dirRight
	dirUp
	dirDown
	dirLeftUp
	dirRightUp
	dirLeftDown
	dirRightDown
)

var allDirs = []dir{dirLeft, dirRight, dirUp, dirDown, dirLeftUp, dirRightUp, dirLeftDown, dirRightDown}

func (d dir) step(x, y int) (int, int) {
	switch d {
	case dirDown:
		return x, y - 1
	case dirLeft:
		return x - 1, y
	case dirRight:
		return x + 1, y
	case dirUp:
		return x, y + 1
	case dirLeftUp:
		return x - 1, y + 1
	case dirLeftDown:
		return x - 1, y - 1
	case dirRightUp:
		return x + 1, y + 1
	case dirRightDown:
		return x + 1, y - 1
	default:
		panic(fmt.Sprintf("unexpected day4.dir: %#v", d))
	}
}

type grid struct {
	lines []string
}

func new(input string) grid {
	lines := strings.Split(input, "\n")
	g := grid{
		lines: lines,
	}
	return g
}

func (g grid) get(x, y int) (rune, bool) {
	if !g.boundsCheck(x, y) {
		return ' ', false
	} else {
		return rune(g.lines[y][x]), true
	}
}

func (g grid) sub(x, y, len int, d dir) (string, bool) {
	sub := ""
	for range len {
		dr, ok := g.get(x, y)
		if !ok {
			return "", false
		}
		sub += string(dr)
		x, y = d.step(x, y)
	}

	return sub, true
}

func (g grid) boundsCheck(x, y int) bool {
	return y >= 0 && y < len(g.lines) && x >= 0 && x < len(g.lines[y])
}

func (g grid) crossPattern(pattern string, mid rune) int {
	if len(pattern)%2 == 0 {
		return -1
	}

	matches := 0
	reverse := aoc.Reverse(pattern)
	for y, line := range g.lines {
		for x, r := range line {
			if mid != r {
				continue
			}

			dx, dy := dirLeftDown.step(x, y)
			down, ok := g.sub(dx, dy, len(pattern), dirRightUp)
			if !ok {
				continue
			}
			if down != pattern && down != reverse {
				continue
			}

			dx, dy = dirLeftUp.step(x, y)
			up, ok := g.sub(dx, dy, len(pattern), dirRightDown)
			if !ok {
				continue
			}

			if up != pattern && up != reverse {
				continue
			}

			matches++
		}
	}
	return matches
}

func (g grid) countPattern(pattern string) int {
	matches := 0
	for y, line := range g.lines {
		for x := range line {
			if g.lines[y][x] != pattern[0] {
				continue
			}

			for _, d := range allDirs {
				sub, ok := g.sub(x, y, len(pattern), d)
				if ok && sub == pattern {
					matches++
				}
			}
		}
	}
	return matches
}
