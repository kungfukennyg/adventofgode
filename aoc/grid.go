package aoc

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

type Grid[T any] struct {
	points [][]T
}

func NewGrid[T any](height, width int) *Grid[T] {
	g := Grid[T]{}
	for range height {
		g.AddRow(make([]T, width))
	}
	return &g
}

func (g *Grid[T]) Len() int {
	if len(g.points) == 0 {
		return 0
	}

	return len(g.points) * len(g.points[0])
}

func (g *Grid[T]) Height() int {
	return len(g.points)
}

func (g *Grid[T]) Width() int {
	if len(g.points) < 1 {
		return 0
	}

	return len(g.points[0])
}

func (g *Grid[T]) Set(p Pos, t T) {
	g.points[p.Y][p.X] = t
}

func (g *Grid[T]) Get(p Pos) T {
	return g.points[p.Y][p.X]
}

func (g *Grid[T]) SetByXY(x, y int, t T) {
	g.points[y][x] = t
}

func (g *Grid[T]) BoundsCheck(p Pos) bool {
	return p.Y >= 0 && p.Y < len(g.points) &&
		p.X >= 0 && p.X < len(g.points[p.Y])
}

func (g *Grid[T]) AddRow(row []T) {
	g.points = append(g.points, row)
}

func (g *Grid[T]) GetRow(y int) []T {
	return g.points[y]
}

// DFS performas a depth-first search of the grid, starting at start and
// ending using the provided function
func (g *Grid[T]) DFS(start Pos,
	neighborFn func(cur Pos) []Pos,
	keepFn func(cur Pos) bool) []Pos {

	if !g.BoundsCheck(start) {
		return []Pos{}
	}

	visited := Set[Pos]{}
	var backtrack func(Pos) []Pos
	backtrack = func(cur Pos) []Pos {
		if !visited.Add(cur) {
			return nil
		}

		ret := []Pos{}
		if keepFn(cur) {
			ret = append(ret, cur)
		}

		nbs := neighborFn(cur)
		for _, nb := range nbs {
			ret = append(ret, backtrack(nb)...)
		}

		return ret
	}

	found := backtrack(start)
	return found
}

func (g *Grid[T]) Points() iter.Seq[Pos] {
	return func(yield func(Pos) bool) {
		for y, row := range g.points {
			for x := range row {
				if !yield(Pos{x, y}) {
					return
				}
			}
		}
	}
}

func (g *Grid[T]) Values() iter.Seq2[Pos, T] {
	return func(yield func(Pos, T) bool) {
		for y, row := range g.points {
			for x, t := range row {
				if !yield(Pos{x, y}, t) {
					return
				}
			}
		}
	}
}

func (g *Grid[T]) Neighbors(p Pos, dirs []Dir) iter.Seq2[Pos, T] {
	return func(yield func(Pos, T) bool) {
		for _, d := range dirs {
			pos := p.Step(d)
			if !g.BoundsCheck(pos) {
				continue
			}

			if !yield(pos, g.Get(pos)) {
				return
			}
		}
	}
}

func (g *Grid[T]) NeighborVecs(p Pos, dirs []Dir) iter.Seq2[Vec, T] {
	return func(yield func(Vec, T) bool) {
		for _, d := range dirs {
			pos := p.Step(d)
			if !g.BoundsCheck(pos) {
				continue
			}

			if !yield(Vec{P: pos, D: d}, g.Get(pos)) {
				return
			}
		}
	}
}

func (g *Grid[T]) Rows() iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		for y, row := range g.points {
			if !yield(y, row) {
				break
			}
		}
	}
}

func (g *Grid[T]) String() string {
	var sb strings.Builder
	y := 0
	for p, r := range g.Values() {
		if p.Y != y {
			y = p.Y
			sb.WriteString("\n")
		}

		switch v := any(r).(type) {
		case rune:
			sb.WriteRune(v)
		case int:
			sb.WriteString(strconv.Itoa(v))
		case string:
			sb.WriteString(v)
		default:
			sb.WriteString(fmt.Sprintf("%v", v))
		}
	}

	return sb.String()
}
