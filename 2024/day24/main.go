package day24

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
	"github.com/kungfukennyg/adventofgode/aoc/opt"
)

type op string

const (
	opAnd op = "AND"
	opXor op = "XOR"
	opOr  op = "OR"
)

func (o op) exec(a, b bool) bool {
	switch o {
	case opAnd:
		return a && b
	case opOr:
		return a || b
	case opXor:
		return a != b
	default:
		panic(fmt.Sprintf("op.exec: unexpected op: %#v", o))
	}
}

type cpu map[string]*wire

type wire struct {
	read func(c cpu) bool
	a, b string
	op   op
}

func (c cpu) read(wire string) bool {
	return c[wire].read(c)
}

func (wires cpu) readPrefixed(target string) int64 {
	var bf aoc.Bitflag
	for r, w := range wires {
		if !strings.HasPrefix(r, target) {
			continue
		}

		bit := aoc.MustAtoi(r[1:])
		v := w.read(wires)
		bf = bf.Set(bit, v)
	}
	return int64(bf)
}

func (c cpu) findWire(a, b opt.Some[string], o op) opt.Some[string] {
	if !(a.Ok() && b.Ok()) {
		return opt.None[string]()
	}

	wa, wb := a.Get(), b.Get()
	for dst, wire := range c {
		if wire.op == o &&
			((wire.a == wa && wire.b == wb) ||
				(wire.a == wb && wire.b == wa)) {
			return opt.From(dst)
		}
	}
	return opt.None[string]()
}

func (c cpu) swapWires(a, b opt.Some[string]) {
	wa, wb := a.Get(), b.Get()
	c[wa], c[wb] = c[wb], c[wa]
}

func findCrossed(input, target string) string {
	c := parse(input)
	bits := 0
	for dst := range c {
		if strings.HasPrefix(dst, "x") {
			bits++
		}
	}

	crossed := []string{}
	carry := opt.None[string]()
	targetWire := func(s string) bool { return strings.HasPrefix(s, target) }
	cross := func(a, b opt.Some[string]) {
		crossed = append(crossed, a.Get(), b.Get())
		c.swapWires(a, b)
	}

	for bit := range bits {
		x := opt.From(fmt.Sprintf("x%02d", bit))
		y := opt.From(fmt.Sprintf("y%02d", bit))

		adder, next := c.findWire(x, y, opXor), c.findWire(x, y, opAnd)
		dst, nextCarry := opt.None[string](), opt.None[string]()

		if carry.Ok() {
			res := c.findWire(adder, carry, opAnd)
			if !res.Ok() {
				cross(adder, next)
				adder, next = next, adder
				res = c.findWire(carry, adder, opAnd)
			}

			dst = c.findWire(adder, carry, opXor)

			for _, w := range []*opt.Some[string]{&adder, &next, &res} {
				if w.Is(targetWire) {
					cross(*w, dst)
					*w, dst = dst, *w
				}
			}

			nextCarry = c.findWire(next, res, opOr)
		}

		if bit != (bits-1) && nextCarry.Is(targetWire) {
			cross(nextCarry, dst)
			nextCarry, dst = dst, nextCarry
		}

		if carry.Ok() {
			carry = nextCarry
		} else {
			carry = next
		}
	}

	slices.Sort(crossed)
	return strings.Join(crossed, ",")
}

func run(input, target string) int64 {
	c := parse(input)
	return c.readPrefixed(target)
}

func parse(input string) cpu {
	wires := cpu{}

	for _, line := range aoc.Lines(input) {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ": ")

			reg := parts[0]
			v := aoc.MustBool(parts[1])
			wires[reg] = &wire{
				read: func(cpu) bool {
					return v
				},
				a: reg,
			}
		} else if len(strings.TrimSpace(line)) > 0 {
			parts := strings.Split(line, " ")
			a, b, dst := parts[0], parts[2], parts[4]
			o := op(parts[1])
			wires[dst] = &wire{
				read: func(c cpu) bool {
					return o.exec(c.read(a), c.read(b))
				},
				a:  a,
				b:  b,
				op: o,
			}
		}
	}

	return wires
}
