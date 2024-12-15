package aoc

import (
	"fmt"
	"iter"
)

type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft

	DirUpRight
	DirUpLeft
	DirDownRight
	DirDownLeft
)

func (d Dir) Step() (int, int) {
	switch d {
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	case DirRight:
		return 1, 0
	case DirUp:
		return 0, -1
	case DirDownRight:
		return 1, 1
	case DirUpLeft:
		return -1, -1
	case DirUpRight:
		return 1, -1
	case DirDownLeft:
		return -1, 1
	default:
		panic(fmt.Sprintf("unexpected aoc.Dir: %#v", d))
	}
}

func (d Dir) Decompose() []Dir {
	switch d {
	case DirDownLeft:
		return []Dir{DirDown, DirLeft}
	case DirDownRight:
		return []Dir{DirDown, DirRight}
	case DirUpLeft:
		return []Dir{DirUp, DirLeft}
	case DirUpRight:
		return []Dir{DirUp, DirRight}
	default:
		return []Dir{d}
	}
}

type Pos struct {
	X, Y int
}

func (p Pos) Step(d Dir) Pos {
	dx, dy := d.Step()
	return Pos{p.X + dx, p.Y + dy}
}

func (p Pos) Sub(b Pos) Pos {
	return Pos{p.X - b.X, p.Y - b.Y}
}

type Grid[T any] struct {
	points [][]T
}

func (g *Grid[T]) SetByPos(t T, p Pos) {
	g.points[p.Y][p.X] = t
}

func (g *Grid[T]) GetByPos(p Pos) T {
	return g.points[p.Y][p.X]
}

func (g *Grid[T]) Get(x, y int) T {
	return g.points[y][x]
}

func (g *Grid[T]) GetInBounds(p Pos) *T {
	if !g.BoundsCheck(p) {
		return nil
	}

	t := g.GetByPos(p)
	return &t
}

func (g *Grid[T]) Set(x, y int, t T) {
	g.points[y][x] = t
}

func (g *Grid[T]) BoundsCheck(p Pos) bool {
	return p.Y >= 0 && p.Y < len(g.points) && p.X >= 0 && p.X < len(g.points[p.Y])
}

func (g *Grid[T]) AddRow(row []T) {
	g.points = append(g.points, row)
}

func (g *Grid[T]) Points() iter.Seq2[Pos, T] {
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

			if !yield(pos, g.GetByPos(pos)) {
				return
			}
		}
	}
}
