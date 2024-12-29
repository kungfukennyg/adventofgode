package day20

import (
	"github.com/kungfukennyg/adventofgode/aoc"
)

type board struct {
	grid       aoc.Grid[rune]
	start, end aoc.Pos
	cheating   bool
}

// Cost implements aoc.Pather.Cost by assigning steps and turns different costs.
func (g board) Cost(from, to aoc.Vec) int {
	return from.P.ManhattanDistance(to.P)
}

// Heuristic implements aoc.Pather.Heuristic by using Manhattan distance.
func (g board) Heuristic(from, to aoc.Vec) int {
	return from.P.ManhattanDistance(to.P)
}

// Neighbors implements aoc.Pather.Neighbors by returning tiles in the current
// direction, as well as tiles on a 90/-90 degree turn.
func (b board) Neighbors(v aoc.Vec) []aoc.Vec {
	nbs := []aoc.Vec{}
	for p, r := range b.grid.NeighborVecs(v.P, aoc.CardinalDirs) {
		if b.cheating || r != '#' {
			nbs = append(nbs, p)
		}
	}
	return nbs
}

// Goal implements aoc.Pather.Goal by returning when we reach the end position.
func (b board) Goal(v, end aoc.Vec) bool {
	return v.P == end.P
}

type cheat struct {
	start, end aoc.Pos
}

func (b board) reachableCheats(cur aoc.Pos, depth int) map[cheat]int {
	b.cheating = true
	defer func() { b.cheating = false }()

	cheats := map[cheat]int{}
	for p, r := range b.grid.Values() {
		dist := cur.ManhattanDistance(p)
		if dist > depth {
			continue
		}
		if r == '#' {
			continue
		}
		ch := cheat{start: cur, end: p}
		if _, ok := cheats[ch]; ok {
			continue
		}

		path := aoc.ShortestPath(b, aoc.Vec{P: cur}, aoc.Vec{P: p})
		for _, pth := range path {
			// need at least 1 wall to pass through to count as a cheat
			if b.grid.Get(pth.P) == '#' {
				if former, ok := cheats[ch]; !ok {
					cheats[ch] = len(path)
				} else {
					cheats[ch] = min(former, len(path))
				}

				break
			}
		}
	}
	return cheats
}

func cheats(input string, cheatDepth, desiredDelta int) int {
	b := parse(input)
	start, end := aoc.Vec{P: b.start}, aoc.Vec{P: b.end}

	basePath := aoc.ShortestPath(b, start, end)
	posToEndCosts := map[aoc.Pos]int{}
	for i, p := range basePath {
		posToEndCosts[p.P] = len(basePath[i:])
	}

	deltas := map[cheat]int{}

	for i, cur := range basePath {
		cheats := b.reachableCheats(cur.P, cheatDepth)
		for cheat, cheatPath := range cheats {
			if _, ok := deltas[cheat]; ok {
				continue
			}

			var remainder int
			if rmd, ok := posToEndCosts[cheat.end]; ok {
				remainder = rmd
			} else {
				remainder = len(aoc.ShortestPath(b, aoc.Vec{P: cheat.end}, end))
				posToEndCosts[cheat.end] = remainder
			}

			delta := len(basePath) - (i + cheatPath + remainder - 1)
			if delta >= desiredDelta {
				deltas[cheat] = delta
			}
		}
	}

	return len(deltas)
}

func (b board) drawPath(cheat, endCheat aoc.Vec, path []aoc.Vec) string {
	for _, p := range path {
		var r rune
		switch p.D {
		case aoc.DirUp:
			r = '^'
		case aoc.DirRight:
			r = '>'
		case aoc.DirDown:
			r = 'v'
		case aoc.DirLeft:
			r = '<'
		default:
			r = 'O'
		}

		b.grid.Set(p.P, r)
	}

	b.grid.Set(b.start, 'S')
	b.grid.Set(cheat.P, '1')
	b.grid.Set(endCheat.P, '2')
	b.grid.Set(b.end, 'E')

	return b.grid.String()
}

func parse(input string) board {
	b := board{
		grid: aoc.Grid[rune]{},
	}

	for y, line := range aoc.Lines(input) {
		b.grid.AddRow(make([]rune, len(line)))
		for x, r := range line {
			switch r {
			case 'S':
				b.start = aoc.Pos{X: x, Y: y}
			case 'E':
				b.end = aoc.Pos{X: x, Y: y}
			}
			b.grid.SetByXY(x, y, r)
		}
	}
	return b
}
