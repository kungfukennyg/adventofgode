package day16

import (
	"github.com/kungfukennyg/adventofgode/aoc"
)

const stepCost = 1
const turnCost = 1000
const wall, empty = '#', '.'

type board struct {
	g          aoc.Grid[rune]
	start, end aoc.Vec
}

// Cost implements aoc.Pather.Cost by assigning steps and turns different costs.
func (g *board) Cost(from, to aoc.Vec) int {
	if to.D == from.D.Clockwise() || to.D == from.D.Counterclockwise() {
		return turnCost
	}

	if from.P.ManhattanDistance(to.P) == 1 {
		return stepCost
	}

	return 0
}

// Heuristic implements aoc.Pather.Heuristic by using Manhattan distance.
func (g *board) Heuristic(from, to aoc.Vec) int {
	dist := from.P.ManhattanDistance(to.P)
	if from.P.X != to.P.X || from.P.Y != to.P.Y {
		dist += turnCost
	}

	return dist
}

// Neighbors implements aoc.Pather.Neighbors by returning tiles in the current
// direction, as well as tiles on a 90/-90 degree turn.
func (b *board) Neighbors(v aoc.Vec) []aoc.Vec {
	o := []aoc.Vec{}
	dirs := []aoc.Dir{v.D, v.D.Clockwise(), v.D.Counterclockwise()}
	for _, d := range dirs {
		p := v.P.Step(d)
		if !b.g.BoundsCheck(p) || b.g.Get(p) == wall {
			continue
		}
		o = append(o, aoc.Vec{P: p, D: d})
	}
	return o
}

// Goal implements aoc.Pather.Goal by returning when we reach the end position. The end direction
// is included for equality comparison if not set to aoc.DirNone.
func (b *board) Goal(v, end aoc.Vec) bool {
	if end.D != aoc.DirNone {
		return end == v
	}

	return end.P == v.P
}

// cost calculates the cost of a path.
func (b *board) cost(start aoc.Vec, path []aoc.Vec) int {
	if len(path) == 0 {
		return 0
	}
	cur := start
	steps, turns := 0, 0
	for _, p := range path {
		steps += 1
		if p.D != cur.D {
			turns++
		}
		cur = p
	}

	if path[0] == start {
		steps--
	}

	return steps + (turnCost * turns)
}

// pathCost returns the shortest path from start to end, and the cost of that path.
func (b *board) pathCost(start aoc.Vec, end aoc.Vec) ([]aoc.Vec, int) {
	path := aoc.ShortestPath(b, start, end)
	cost := b.cost(start, path)
	return path, cost
}

// sharedPaths returns the number of positions shared amongst the lowest-cost
// paths from the start to end.
func (b *board) sharedPaths() int {
	bestPath, lowestCost := b.pathCost(b.start, b.end)
	forks := aoc.Set[aoc.Pos]{}
	dirs := aoc.CardinalDirs

	// find all spots with multiple turn options
	for _, v := range bestPath {
		nbs := make([]aoc.Pos, 0, 4)
		for p, r := range b.g.Neighbors(v.P, dirs) {
			if r == empty {
				nbs = append(nbs, p)
			}
		}
		if len(nbs) >= 3 {
			forks.AddAll(nbs)
		}
	}

	shared := aoc.Set[aoc.Pos]{}

	// explore all other possible paths from forks
	for fork := range forks {
		if shared.Contains(fork) {
			continue
		}
		if fork == b.end.P || fork == b.start.P || b.g.Get(fork) == wall {
			continue
		}

		for _, d := range dirs {
			p := fork.Step(d)
			if !b.g.BoundsCheck(p) || b.g.Get(p) == wall {
				continue
			}

			// start->fork
			v := p.WithDir(d)
			pathToStart, cost := b.pathCost(b.start, v)
			if cost+int(b.Heuristic(v, b.end)) > lowestCost {
				break
			}

			// fork->end
			pathToEnd, costEnd := b.pathCost(v, b.end)
			if cost+costEnd == lowestCost {
				for _, s := range pathToStart {
					shared.Add(s.P)
				}
				for _, s := range pathToEnd {
					shared.Add(s.P)
				}
			}
		}
	}

	return len(shared)
}

func shortestPathCost(input string) int {
	b := parse(input)
	_, cost := b.pathCost(b.start, b.end)
	return cost
}

func sharedShortestPath(input string) int {
	b := parse(input)
	return b.sharedPaths()
}

func parse(input string) *board {
	b := &board{g: aoc.Grid[rune]{}}
	for y, line := range aoc.Lines(input) {
		b.g.AddRow(make([]rune, len(line)))
		for x, r := range line {
			switch r {
			case 'S':
				b.start = aoc.NewVec(x, y, aoc.DirRight)
			case 'E':
				b.end = aoc.NewVec(x, y, aoc.DirNone)
			}
			b.g.SetByXY(x, y, r)
		}
	}
	return b
}
