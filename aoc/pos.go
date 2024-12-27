package aoc

import "fmt"

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

func (p Pos) Add(b Pos) Pos {
	return Pos{p.X + b.X, p.Y + b.Y}
}

func (p Pos) AddXY(x, y int) Pos {
	return Pos{p.X + x, p.Y + y}
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
