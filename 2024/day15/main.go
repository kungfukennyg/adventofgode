package day15

import (
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type board struct {
	grid  *aoc.Grid[rune]
	robot aoc.Pos
	move  aoc.Dir
	wide  bool
}

func (b *board) isBox(r rune) bool {
	if b.wide {
		return r == '[' || r == ']'
	}
	return r == 'O'
}

func (b *board) neighbors(cur aoc.Pos) []aoc.Pos {
	next := cur.Step(b.move)
	if !b.grid.BoundsCheck(next) ||
		(b.grid.Get(next) == '#' || b.grid.Get(next) == '.') {
		return []aoc.Pos{}
	}

	pos := []aoc.Pos{next}
	if !b.wide || (b.move == aoc.DirLeft || b.move == aoc.DirRight) {
		return pos
	}

	r := b.grid.Get(next)
	if r == '[' {
		pos = append(pos, next.Step(aoc.DirRight))
	} else if r == ']' {
		pos = append(pos, next.Step(aoc.DirLeft))
	}

	return pos
}

func (b *board) moveBoxes(boxes []aoc.Pos) {
	moved := make([]rune, len(boxes))
	pos := aoc.Set[aoc.Pos]{}
	for i, box := range boxes {
		moved[i] = b.grid.Get(box)
		pos.Add(box.Step(b.move))
	}

	for i, box := range boxes {
		next := box.Step(b.move)
		b.grid.Set(next, moved[i])
		if !pos.Contains(box) {
			b.grid.Set(box, '.')
		}
	}

	next := b.robot.Step(b.move)
	b.grid.Set(next, '@')
	b.grid.Set(b.robot, '.')
	b.robot = next
}

func (b *board) step() {
	next := b.robot.Step(b.move)
	r := b.grid.Get(next)
	if r == '.' {
		b.grid.Set(b.robot, '.')
		b.grid.Set(next, '@')
		b.robot = next
		return
	} else if r == '#' {
		return
	} else if b.isBox(r) {
		boxes := b.grid.DFS(b.robot, b.neighbors,
			func(cur aoc.Pos) bool { return b.isBox(b.grid.Get(cur)) })

		canMove := true
		for _, box := range boxes {
			next := box.Step(b.move)
			if b.grid.Get(next) == '#' {
				canMove = false
				break
			}
		}

		if canMove {
			b.moveBoxes(boxes)
		}
	}
}

func simulate(input string, wide bool) int {
	b, moves := parse(input, wide)
	for _, move := range moves {
		b.move = move
		b.step()
	}

	sum := 0
	for p, r := range b.grid.Values() {
		if r != 'O' && r != '[' {
			continue
		}

		sum += (100 * p.Y) + p.X
	}
	return sum
}

func parse(input string, wide bool) (*board, []aoc.Dir) {
	b := &board{
		grid: &aoc.Grid[rune]{},
		wide: wide,
	}

	for y, line := range aoc.Lines(input) {
		if strings.ContainsAny(line, "<v>^") || len(line) == 0 {
			continue
		}

		cur := []rune{}
		for _, s := range line {
			switch s {
			case '@':
				b.robot = aoc.Pos{X: len(cur), Y: y}
				cur = append(cur, s)
				if wide {
					cur = append(cur, '.')
				}
			case 'O':
				if wide {
					cur = append(cur, '[', ']')
				} else {
					cur = append(cur, s)
				}
			default:
				cur = append(cur, s)
				if wide {
					cur = append(cur, s)
				}
			}
		}

		b.grid.AddRow(cur)

	}

	moves := parseMoves(input)
	return b, moves
}

func parseMoves(input string) []aoc.Dir {
	moves := []aoc.Dir{}
	for _, line := range aoc.Lines(input) {
		if !strings.ContainsAny(line, "<v>^") {
			continue
		}

		var move aoc.Dir
		for _, s := range line {
			switch s {
			case '<':
				move = aoc.DirLeft
			case '^':
				move = aoc.DirUp
			case '>':
				move = aoc.DirRight
			case 'v':
				move = aoc.DirDown
			}
			moves = append(moves, move)
		}
	}
	return moves
}
