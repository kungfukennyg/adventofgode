package day1

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type move struct {
	right bool
	steps int
}

func firstPlaceVisitedTwice(input string) int {
	moves := parseMoves(input)

	start := aoc.Vec{
		D: aoc.DirUp,
	}
	path := tracePath(moves, start)
	visited := aoc.Set[aoc.Pos]{}
	var goal aoc.Pos
	for _, p := range path {
		if !visited.Add(p) {
			goal = p
			break
		}
	}

	return goal.ManhattanDistance(start.P)
}

func hqDistance(input string) int {
	moves := parseMoves(input)

	start := aoc.Vec{
		D: aoc.DirUp,
	}
	path := tracePath(moves, start)
	goal := path[len(path)-1]

	dist := goal.ManhattanDistance(start.P)
	return dist
}

func tracePath(moves []move, start aoc.Vec) []aoc.Pos {
	path := []aoc.Pos{}
	cur := start
	for _, m := range moves {
		if m.right {
			cur.D = cur.D.Clockwise()
		} else {
			cur.D = cur.D.Counterclockwise()
		}

		for range m.steps {
			cur.P = cur.P.Step(cur.D)
			path = append(path, cur.P)
		}
	}

	return path
}

func parseMoves(input string) []move {
	moves := []move{}
	for _, part := range strings.Split(input, ", ") {
		m := move{}
		switch part[0] {
		case 'R':
			m.right = true
		}

		m.steps = aoc.MustAtoi(part[1:])
		moves = append(moves, m)
	}

	return moves
}
