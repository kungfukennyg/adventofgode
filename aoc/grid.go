package aoc

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

type Dir int

const (
	DirNone Dir = iota
	DirUp
	DirRight
	DirDown
	DirLeft

	DirUpRight
	DirUpLeft
	DirDownRight
	DirDownLeft
)

var CardinalDirs = []Dir{DirRight, DirDown, DirLeft, DirUp}

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

func (d Dir) Clockwise() Dir {
	switch d {
	case DirDown:
		return DirLeft
	case DirLeft:
		return DirUp
	case DirUp:
		return DirRight
	case DirRight:
		return DirDown
	default:
		panic(fmt.Sprintf("unexpected aoc.Dir: %#v", d))
	}
}

func (d Dir) Counterclockwise() Dir {
	switch d {
	case DirDown:
		return DirRight
	case DirRight:
		return DirUp
	case DirUp:
		return DirLeft
	case DirLeft:
		return DirDown
	default:
		panic(fmt.Sprintf("unexpected aoc.Dir: %#v", d))
	}
}

func (p Pos) ManhattanDistance(o Pos) int {
	return Abs(p.X-o.X) + Abs(p.Y-o.Y)
}

func (d Dir) String() string {
	switch d {
	case DirDown:
		return "DirDown"
	case DirDownLeft:
		return "DirDownLeft"
	case DirDownRight:
		return "DirDownRight"
	case DirLeft:
		return "DirLeft"
	case DirRight:
		return "DirRight"
	case DirUp:
		return "DirUp"
	case DirUpLeft:
		return "DirUpLeft"
	case DirUpRight:
		return "DirUpRight"
	default:
		panic(fmt.Sprintf("unexpected aoc.Dir: %#v", d))
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

func (p Pos) WithDir(d Dir) Vec {
	return Vec{P: p, D: d}
}

// Vec combines an xy position and a direction.
type Vec struct {
	P Pos
	D Dir
}

func NewVec(x, y int, d Dir) Vec {
	return Vec{P: Pos{X: x, Y: y}, D: d}
}

type Grid[T any] struct {
	points [][]T
}

func (g *Grid[T]) Len() int {
	if len(g.points) == 0 {
		return 0
	}

	return len(g.points) * len(g.points[0])
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
