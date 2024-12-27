package day18

import (
	"log"
	"strconv"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type board struct {
	grid          *aoc.Grid[rune]
	height, width int
}

// Cost supplies the expense for moving between two vertices.
func (b *board) Cost(from, to aoc.Vec) int {
	return from.P.ManhattanDistance(to.P)
}

// Heuristic supplies an estimate of the minimum cost from any vertex to the goal
func (b *board) Heuristic(from, to aoc.Vec) int {
	return from.P.ManhattanDistance(to.P)
}

// Neighbors returns the neighboring vertices to p.
func (b *board) Neighbors(p aoc.Vec) []aoc.Vec {
	neighbors := []aoc.Vec{}
	for o, v := range b.grid.Neighbors(p.P, aoc.CardinalDirs) {
		if v != '#' {
			neighbors = append(neighbors, aoc.Vec{P: o})
		}
	}
	return neighbors
}

// Goal returns whether the current position matches the goal vertex.
func (b *board) Goal(cur, goal aoc.Vec) bool {
	return cur.P == goal.P
}

func (b *board) drawPath(path []aoc.Vec) {
	grid := aoc.NewGrid[rune](b.height, b.width)
	for p, v := range b.grid.Values() {
		grid.Set(p, v)
	}

	for _, p := range path {
		grid.Set(p.P, 'O')
	}

	log.Println("\n", grid.String())
}

func cutsOffExit(input string, height, width int) string {
	b, coords := parse(input, height, width)
	start, goal := aoc.Pos{X: 0, Y: 0}, aoc.Pos{X: width - 1, Y: height - 1}

	x, y := "-1", "-1"
	for _, p := range coords {
		b.grid.Set(p, '#')

		path := aoc.ShortestPath(b, aoc.Vec{P: start}, aoc.Vec{P: goal})
		log.Printf("%v:\n", p)
		b.drawPath(path)
		if len(path) == 0 {
			x, y = strconv.Itoa(p.X), strconv.Itoa(p.Y)
			break
		}
	}

	return strings.Join([]string{x, y}, ",")
}

func findExit(input string, height, width, bytes int) int {
	b, coords := parse(input, height, width)
	for i := range bytes {
		p := coords[i]
		b.grid.Set(p, '#')
	}

	start, goal := aoc.Pos{X: 0, Y: 0}, aoc.Pos{X: width - 1, Y: height - 1}
	path := aoc.ShortestPath(b, aoc.Vec{P: start}, aoc.Vec{P: goal})
	return len(path) - 1
}

func parse(input string, height, width int) (*board, []aoc.Pos) {
	grid := aoc.NewGrid[rune](height, width)
	for p := range grid.Points() {
		grid.Set(p, '.')
	}
	coords := []aoc.Pos{}
	for _, line := range aoc.Lines(input) {
		xy := aoc.Ints(line, ",")
		coords = append(coords, aoc.Pos{X: xy[0], Y: xy[1]})
	}

	return &board{grid: grid, height: height, width: width}, coords
}
