package day10

import (
	"fmt"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type dir int

const (
	dirUp dir = iota
	dirRight
	dirDown
	dirLeft
)

func (d dir) step() (int, int) {
	switch d {
	case dirDown:
		return 0, 1
	case dirLeft:
		return -1, 0
	case dirRight:
		return 1, 0
	case dirUp:
		return 0, -1
	default:
		panic(fmt.Sprintf("unexpected day10.dir: %#v", d))
	}
}

type pos struct {
	x, y int
}

func (p pos) step(d dir) pos {
	dx, dy := d.step()
	return pos{p.x + dx, p.y + dy}
}

type grid struct {
	heights    [][]int
	trailHeads []pos
	trailEnds  []pos
}

func (g grid) get(p pos) int {
	return g.heights[p.y][p.x]
}

func (g grid) boundsCheck(p pos) bool {
	return p.y >= 0 && p.y < len(g.heights) && p.x >= 0 && p.x < len(g.heights[p.y])
}

func (g grid) trails(head pos, dirs []dir) (int, int) {
	visited := map[pos]struct{}{}
	total := 0
	var backtrack func(p pos, visited map[pos]struct{})
	backtrack = func(p pos, visited map[pos]struct{}) {
		for _, d := range dirs {
			n := g.get(p)
			p := p.step(d)
			if !g.boundsCheck(p) {
				continue
			}

			if g.get(p) != n+1 {
				continue
			}
			if g.get(p) == 9 {
				total++
				if _, ok := visited[p]; ok {
					continue
				}

				visited[p] = struct{}{}
				continue
			}

			backtrack(p, visited)
		}
	}

	backtrack(head, visited)
	return len(visited), total
}

func scoreTrailheads(input string) (int, int) {
	g := parse(input)
	trailheads, trails := 0, 0
	for _, start := range g.trailHeads {
		unique, total := g.trails(start, []dir{dirLeft, dirDown, dirRight, dirUp})
		trailheads += unique
		trails += total
	}
	return trailheads, trails
}

func parse(input string) grid {
	g := grid{heights: [][]int{}}
	for y, line := range aoc.Lines(input) {
		g.heights = append(g.heights, make([]int, len(line)))
		for x, s := range line {
			n := aoc.MustAtoi(string(s))
			g.heights[y][x] = n
			if n == 0 {
				g.trailHeads = append(g.trailHeads, pos{x, y})
			} else {
				g.trailEnds = append(g.trailEnds, pos{x, y})
			}
		}
	}
	return g
}
