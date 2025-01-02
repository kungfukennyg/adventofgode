package day21

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
	"github.com/kungfukennyg/adventofgode/aoc/list"
)

var rotatedCardinals [][]aoc.Dir

type keypad struct {
	grid      *aoc.Grid[string]
	cursor    aoc.Pos
	codeToPos map[string]aoc.Pos
	dirs      []aoc.Dir
}

// Cost implements aoc.Pather.Cost by assigning steps and turns different costs.
func (k keypad) Cost(from, to aoc.Vec) int {
	// favor direction we're already moving
	cost := from.P.ManhattanDistance(to.P)
	if from.D != aoc.DirNone && to.D != aoc.DirNone {
		if from.D != to.D {
			cost++
		}
	}
	return cost
}

// Heuristic implements aoc.Pather.Heuristic by using Manhattan distance.
func (k keypad) Heuristic(from, to aoc.Vec) int {
	return k.Cost(from, to)
}

// Neighbors implements aoc.Pather.Neighbors by returning tiles in the four
// cardinal directions.
func (k keypad) Neighbors(v aoc.Vec) []aoc.Vec {
	nbs := []aoc.Vec{}
	for p, r := range k.grid.NeighborVecs(v.P, rotatedCardinals[rand.Intn(len(rotatedCardinals))]) {
		if r != "" {
			nbs = append(nbs, p)
		}
	}
	return nbs
}

// Goal implements aoc.Pather.Goal by returning when we reach the end position.
func (k keypad) Goal(v, end aoc.Vec) bool {
	return v.P == end.P
}

func (k keypad) getPos(str string) aoc.Pos {
	if str == "" {
		panic(fmt.Errorf("keypad.getPos: tried to access empty slot '%s'", str))
	}

	return k.codeToPos[str]
}

func newKeypad(layout [][]string) keypad {
	if len(layout) == 0 {
		return keypad{grid: &aoc.Grid[string]{}}
	}

	kp := keypad{
		grid:      aoc.NewGrid[string](len(layout), len(layout[0])),
		cursor:    aoc.Pos{},
		codeToPos: map[string]aoc.Pos{},
		dirs:      aoc.CardinalDirs,
	}
	kp.codeToPos = make(map[string]aoc.Pos, kp.grid.Len())
	for y, line := range layout {
		for x, s := range line {
			p := aoc.Pos{X: x, Y: y}
			kp.grid.Set(p, s)
			kp.codeToPos[s] = p
		}
	}
	return kp
}

func dirToButton(d aoc.Dir) string {
	switch d {
	case aoc.DirUp:
		return "^"
	case aoc.DirRight:
		return ">"
	case aoc.DirDown:
		return "v"
	case aoc.DirLeft:
		return "<"
	default:
		panic(fmt.Errorf("dirToButton: unexpected dir: '%+v'", d))
	}
}

func pathToStr(path []aoc.Vec) string {
	var moves strings.Builder
	for _, step := range path {
		moves.WriteString(dirToButton(step.D))
	}
	return moves.String()
}

func (k *keypad) apply(in string) string {
	translated := ""
	for _, c := range in {
		s := string(c)
		pos := k.getPos(s)
		path := aoc.ShortestPath(k, aoc.Vec{P: k.cursor}, aoc.Vec{P: pos})[1:]
		translated += pathToStr(path) + "A"
		k.cursor = pos
	}
	return translated
}

func evalCode(keypads []*keypad, code []*list.Linked[string]) {
	for _, kp := range keypads {
		for _, part := range code {
			t := part.Tail()
			o := kp.apply(t)
			part.Add(o)
		}
	}
}

type codeSequence struct {
	code  string
	parts []*list.Linked[string]
}

func (cs codeSequence) output() []string {
	out := make([]string, len(cs.parts))
	for _, seq := range cs.parts {
		for i, p := range seq.Iter() {
			out[i] += p
		}
	}
 
	return out
}

func (cs codeSequence) String() string {
	var sb strings.Builder
	for _, out := range cs.output() {
		sb.WriteString(out)
		sb.WriteString("\n")
	}

	for _, seq := range cs.parts {
		sb.WriteString(seq.String())
		sb.WriteString("\n")
	}

	return sb.String()
}

func eval(input string, depth int) int {
	cardinals := slices.Clone(aoc.CardinalDirs)
	rotatedCardinals = make([][]aoc.Dir, len(cardinals))
	for i := range len(rotatedCardinals) {
		rotatedCardinals[i] = aoc.Rotate(cardinals, i)
	}

	lowest := map[string]int{}
	for i := 0; i < 25; i++ {
		keypads, codes := parse(input, depth)
		for _, codeSeq := range codes {
			evalCode(keypads, codeSeq.parts)
			steps := codeSeq.output()
			final := steps[len(steps)-1]

			numericCode := aoc.MustAtoi(strings.ReplaceAll(codeSeq.code, "A", ""))
			len := len(final)
			complexity := numericCode * len
			if cmp, ok := lowest[codeSeq.code]; ok {
				lowest[codeSeq.code] = min(cmp, complexity)
			} else {
				lowest[codeSeq.code] = complexity
			}
		}
	}

	complexity := 0
	for _, cmp := range lowest {
		complexity += cmp
	}
	return complexity
}

func parse(input string, depth int) ([]*keypad, []codeSequence) {
	var num = [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"", "0", "A"},
	}

	var dir = [][]string{
		{"", "^", "A"},
		{"<", "v", ">"},
	}

	keypads := make([]*keypad, 0, depth)
	for i := range depth {
		var kp keypad
		if i == 0 {
			kp = newKeypad(num)
			kp.cursor = aoc.Pos{X: 2, Y: 3}
		} else {
			kp = newKeypad(dir)
			kp.cursor = aoc.Pos{X: 2, Y: 0}
		}
		keypads = append(keypads, &kp)
	}

	moves := []codeSequence{}
	for _, code := range aoc.Lines(input) {
		seq := []*list.Linked[string]{}
		for _, c := range code {
			seq = append(seq, list.NewLinked(string(c)))
		}
		moves = append(moves, codeSequence{
			code:  code,
			parts: seq,
		})
	}
	return keypads, moves
}
