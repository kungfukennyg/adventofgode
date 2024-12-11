package day6

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type op string

const (
	opOn     op = "turn on"
	opOff    op = "turn off"
	opToggle op = "toggle"
)

func opFromStr(s string) op {
	o := op(s)
	switch o {
	case opOff:
	case opOn:
	case opToggle:
	default:
		panic(fmt.Sprintf("unexpected day6.op '%s'", s))
	}
	return o
}

type pos struct {
	x, y int
}

func posFromStr(s string) pos {
	parts := strings.Split(s, ",")
	return pos{
		x: aoc.MustAtoi(parts[0]),
		y: aoc.MustAtoi(parts[1]),
	}
}

type instruction struct {
	op         op
	start, end pos
}

func parse(line string) instruction {
	// turn on 0,0 through 999,999
	buf := bufio.NewScanner(strings.NewReader(line))
	buf.Split(bufio.ScanWords)
	op := aoc.MustText(buf)
	if op == "turn" {
		op += " " + aoc.MustText(buf)
	}

	instr := instruction{op: opFromStr(op)}
	instr.start = posFromStr(aoc.MustText(buf))
	_ = aoc.MustText(buf)
	instr.end = posFromStr(aoc.MustText(buf))
	return instr
}

const SIZE int = 1000

type grid struct {
	lights         [SIZE][SIZE]int
	brightness     int
	brightnessFunc func(o op, b int) int
}

func lightLevel(input string, binary bool) int {
	g := grid{lights: [1000][1000]int{}}
	if binary {
		g.brightnessFunc = g.binary
	} else {
		g.brightnessFunc = g.variable
	}
	for _, line := range strings.Split(input, "\n") {
		instr := parse(line)
		for y := instr.start.y; y <= instr.end.y; y++ {
			for x := instr.start.x; x <= instr.end.x; x++ {
				b := g.lights[x][y]
				updated := g.brightnessFunc(instr.op, b)
				g.brightness += updated - b
				g.lights[x][y] = updated
			}
		}
	}
	return g.brightness
}

func (g *grid) binary(o op, b int) int {
	switch o {
	case opOff:
		return 0
	case opOn:
		return 1
	case opToggle:
		if b == 0 {
			return 1
		} else {
			return 0
		}
	default:
		panic(fmt.Sprintf("unexpected day6.op: %#v", o))
	}
}

func (g *grid) variable(o op, b int) int {
	switch o {
	case opOff:
		updated := b - 1
		if updated < 0 {
			updated = 0
		}
		return updated
	case opOn:
		return b + 1
	case opToggle:
		return b + 2
	default:
		panic(fmt.Sprintf("unexpected day6.op: %#v", o))
	}
}
