package day14

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type board struct {
	grid   aoc.Grid[int]
	robots []*robot
}

type robot struct {
	pos aoc.Pos
	vel aoc.Pos
}

type quadrants struct {
	ul, ur, dl, dr int
}

func (q quadrants) mult() int {
	return q.ul * q.ur * q.dl * q.dr
}

func (b *board) step() {
	for _, r := range b.robots {
		b.grid.Set(r.pos, b.grid.Get(r.pos)-1)
		r.pos = r.pos.Add(r.vel)
		if r.pos.Y >= b.grid.Height() {
			r.pos.Y -= b.grid.Height()
		} else if r.pos.Y < 0 {
			r.pos.Y = b.grid.Height() + r.pos.Y
		}

		if r.pos.X >= b.grid.Width() {
			r.pos.X -= b.grid.Width()
		} else if r.pos.X < 0 {
			r.pos.X = b.grid.Width() + r.pos.X
		}

		b.grid.Set(r.pos, b.grid.Get(r.pos)+1)
	}
}

func (b *board) quadrants() quadrants {
	hx, hy := b.grid.Width()/2, b.grid.Height()/2
	ul, ur, dl, dr := 0, 0, 0, 0
	for p, v := range b.grid.Values() {
		if p.X < hx && p.Y < hy {
			ul += v
		} else if p.X > hx && p.Y < hy {
			ur += v
		} else if p.X < hx && p.Y > hy {
			dl += v
		} else if p.X > hx && p.Y > hy {
			dr += v
		}
	}
	return quadrants{ul, ur, dl, dr}
}

func christmasTree(input string, height, width int) int {
	b := parse(input, height, width)
outer:
	for i := 0; ; i++ {
		b.step()
		for _, v := range b.grid.Values() {
			if v > 1 {
				continue outer
			}
		}

		return i
	}
}

func safetyFactor(input string, height, width, seconds int) int {
	b := parse(input, height, width)
	for range seconds {
		b.step()
	}

	return b.quadrants().mult()
}

func parse(input string, height, width int) *board {
	grid := aoc.Grid[int]{}
	for range height {
		grid.AddRow(make([]int, width))
	}

	robots := []*robot{}
	for _, line := range aoc.Lines(input) {
		parts := strings.Split(line, " ")
		r := &robot{
			pos: parsePos(parts[0]),
			vel: parsePos(parts[1]),
		}
		robots = append(robots, r)
		grid.Set(r.pos, grid.Get(r.pos)+1)
	}
	return &board{grid, robots}
}

func parsePos(part string) aoc.Pos {
	pos := strings.TrimPrefix(part, "p=")
	pos = strings.TrimPrefix(pos, "v=")
	xy := strings.Split(pos, ",")
	return aoc.Pos{
		X: aoc.MustAtoi(xy[0]),
		Y: aoc.MustAtoi(xy[1]),
	}
}
