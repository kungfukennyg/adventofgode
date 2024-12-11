package day18

import (
	"slices"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type grid struct {
	lights   [][]bool
	alwaysOn [][2]int
}

func (g *grid) clone() *grid {
	o := &grid{
		lights:   make([][]bool, len(g.lights)),
		alwaysOn: slices.Clone(g.alwaysOn),
	}

	for y, row := range g.lights {
		o.lights[y] = slices.Clone(row)
	}
	return o
}

func (g *grid) get(x, y int) bool {
	for _, coord := range g.alwaysOn {
		if coord[0] == x && coord[1] == y {
			return true
		}
	}

	return g.lights[y][x]
}

func (g *grid) boundsCheck(x, y int) bool {
	return y >= 0 && y < len(g.lights) && x >= 0 && x < len(g.lights[y])
}

func (g *grid) lit() int {
	l := 0
	for y, row := range g.lights {
		for x := range row {
			v := g.get(x, y)
			if v {
				l++
			}
		}
	}
	return l
}

func (g *grid) litNeighbors(x, y int) int {
	lit := 0
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}

			if !g.boundsCheck(x+dx, y+dy) {
				continue
			}

			if g.get(x+dx, y+dy) {
				lit++
			}
		}
	}
	return lit
}

func (g *grid) step() *grid {
	o := g.clone()
	for y, row := range g.lights {
		for x := range row {
			lit := g.get(x, y)
			neighbors := g.litNeighbors(x, y)
			if lit && neighbors != 2 && neighbors != 3 {
				o.lights[y][x] = false
			} else if !lit && neighbors == 3 {
				o.lights[y][x] = true
			}
		}
	}

	return o
}

func sim(input string, steps int, cornersAlwaysOn bool) int {
	g := parse(input)
	if cornersAlwaysOn {
		n := len(g.lights) - 1
		g.alwaysOn = [][2]int{{0, 0}, {0, n}, {n, 0}, {n, n}}
	}

	for range steps {
		g = g.step()
	}
	return g.lit()
}

func parse(input string) *grid {
	g := &grid{
		lights: [][]bool{},
	}
	for y, line := range aoc.Lines(input) {
		g.lights = append(g.lights, make([]bool, len(line)))
		for x, s := range line {
			if s == '#' {
				g.lights[y][x] = true
			}
		}
	}
	return g
}
