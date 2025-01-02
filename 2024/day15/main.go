package day15

import (
	"log"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type board struct {
	grid  aoc.Grid[rune]
	robot aoc.Pos
}

func (b *board) move(cur, next aoc.Pos, r rune) {
	b.grid.Set(cur, '.')
	b.grid.Set(next, '@')
	b.robot = next
}

func (b *board) boxesInLine(d aoc.Dir, obj aoc.Pos) aoc.Stack[aoc.Pos] {
	line := aoc.Stack[aoc.Pos]{}
	cur := obj
	for {
		cur = cur.Step(d)
		if !b.grid.BoundsCheck(cur) {
			break
		}
		r := b.grid.Get(cur)
		line.Push(cur)
		if r == '#' || r == '.' {
			break
		}
	}

	return line
}

func (b *board) step(d aoc.Dir, obj aoc.Pos) (aoc.Pos, bool) {
	line := b.boxesInLine(d, obj)
	if line.IsEmpty() {
		return aoc.Pos{}, false
	}

	var cur *aoc.Pos
	for !line.IsEmpty() {
		next, _ := line.Pop()
		if cur == nil {
			cur = &next
			continue
		}

		r, nr := b.grid.Get(*cur), b.grid.Get(next)
		b.grid.Set(next, r)
		b.grid.Set(*cur, nr)
		cur = &next
	}

	if cur == nil {
		return aoc.Pos{}, false
	}

	return *cur, true
}

func (b *board) sumBoxCoordinates() int {
	sum := 0
	for p, r := range b.grid.Values() {
		if r == 'O' {
			sum += p.X + (100 * p.Y)
		}
	}
	return sum
}

func scaledCoordinates(input string) int {
	b, moves := parse(input)

	// scale grid 2x horizontally
	robot := aoc.Pos{}
	scaled := aoc.NewGrid[rune](b.grid.Height(), b.grid.Width()*2)
	for y, row := range b.grid.Rows() {
		dx := 0
		for _, r := range row {
			var lh rune
			var rh rune
			switch r {
			case 'O':
				lh, rh = '[', ']'
			case '@':
				lh, rh = r, '.'
			default:
				lh, rh = r, r
			}

			scaled.SetByXY(dx, y, lh)
			scaled.SetByXY(dx+1, y, rh)
			if lh == '@' {
				robot.X, robot.Y = dx, y
			}
			dx += 2
		}
	}
	b.grid = *scaled
	b.robot = robot

	for _, d := range moves {
		log.Printf("\n%s: \n%s\n", d, b.grid.String())
		if next, ok := b.step(d, b.robot); ok {
			b.robot = next
		}
	}

	return b.sumBoxCoordinates()
}

func boxCoordinates(input string) int {
	b, moves := parse(input)
	for _, d := range moves {
		log.Printf("\n%s: \n%s\n", d, b.grid.String())
		if next, ok := b.step(d, b.robot); ok {
			b.robot = next
		}
	}
	return b.sumBoxCoordinates()
}

func parse(input string) (*board, []aoc.Dir) {
	moves := []aoc.Dir{}
	board := &board{
		grid: aoc.Grid[rune]{},
	}

	for y, line := range aoc.Lines(input) {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if !strings.ContainsAny(line, "<^>v") {
			row := []rune{}
			for x, r := range line {
				switch r {
				case '@':
					board.robot = aoc.Pos{X: x, Y: y}
					row = append(row, '@')
				case 'O':
					row = append(row, r)
				case '.':
					row = append(row, r)
				case '#':
					row = append(row, r)
				}
			}
			board.grid.AddRow(row)
		} else {
			for _, r := range line {
				var d aoc.Dir
				switch r {
				case '>':
					d = aoc.DirRight
				case '<':
					d = aoc.DirLeft
				case '^':
					d = aoc.DirUp
				case 'v':
					d = aoc.DirDown
				}
				moves = append(moves, d)
			}
		}
	}

	return board, moves
}
