package day7

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

type cpuV2 struct {
	wires  map[string]wire
	values map[string]uint16
}
type wire struct {
	valueFn func(cpuV2) uint16
}

func newCpu(input string) cpuV2 {
	c := cpuV2{
		wires:  map[string]wire{},
		values: map[string]uint16{},
	}
	instrs := parse(input)
	for _, i := range instrs {
		// 123 -> x
		switch i.op {
		case "->":
			src := i.a
			c.wires[i.dst] = wire{
				valueFn: func(c cpuV2) uint16 {
					return c.read(src)
				},
			}
		// NOT x -> h
		case "NOT":
			src := i.a
			c.wires[i.dst] = wire{
				valueFn: func(c cpuV2) uint16 {
					return c.not(src)
				},
			}
		case "AND":
			a, b := i.a, i.b
			c.wires[i.dst] = wire{
				valueFn: func(c cpuV2) uint16 {
					return c.and(a, b)
				},
			}
		case "OR":
			a, b := i.a, i.b
			c.wires[i.dst] = wire{
				valueFn: func(c cpuV2) uint16 {
					return c.or(a, b)
				},
			}
		case "LSHIFT":
			a, b := i.a, i.b
			c.wires[i.dst] = wire{
				valueFn: func(c cpuV2) uint16 {
					return c.lshift(a, b)
				},
			}
		case "RSHIFT":
			a, b := i.a, i.b
			c.wires[i.dst] = wire{
				valueFn: func(c cpuV2) uint16 {
					return c.rshift(a, b)
				},
			}
		}
	}
	return c
}

func (c cpuV2) override(src string, value uint16) {
	w := c.wires[src]
	w.valueFn = func(cv cpuV2) uint16 { return value }
	c.wires[src] = w
	clear(c.values)
}

func (c cpuV2) read(src string) uint16 {
	n, err := strconv.ParseUint(src, 10, 16)
	if err != nil {
		if w, ok := c.wires[src]; ok {
			if _, ok := c.values[src]; !ok {
				c.values[src] = w.valueFn(c)
			}

			return c.values[src]
		}
	}

	return uint16(n)
}

func (c cpuV2) not(src string) uint16 {
	return ^c.read(src)
}

func (c cpuV2) and(a, b string) uint16 {
	return c.read(a) & c.read(b)
}

func (c cpuV2) or(a, b string) uint16 {
	return c.read(a) | c.read(b)
}

func (c cpuV2) lshift(a, b string) uint16 {
	return c.read(a) << c.read(b)
}

func (c cpuV2) rshift(a, b string) uint16 {
	return c.read(a) >> c.read(b)
}

type instr struct {
	op        string
	a, b, dst string
}

func parse(instructions string) []instr {
	lines := strings.Split(instructions, "\n")
	instrs := make([]instr, 0, len(lines))

	for _, line := range lines {
		buf := bufio.NewScanner(strings.NewReader(line))
		buf.Split(bufio.ScanWords)
		op := aoc.MustText(buf)
		switch op {
		case "NOT":
			src := aoc.MustText(buf)
			buf.Scan()
			dst := aoc.MustText(buf)
			instrs = append(instrs, instr{
				op:  op,
				a:   src,
				dst: dst,
			})
		default:
			a := op
			op = aoc.MustText(buf)
			b := aoc.MustText(buf)
			if op == "->" {
				instrs = append(instrs, instr{
					op:  op,
					a:   a,
					dst: b,
				})
				continue
			}

			buf.Scan()
			dst := aoc.MustText(buf)
			switch op {
			case "AND":
			case "OR":
			case "LSHIFT":
			case "RSHIFT":
			default:
				panic(fmt.Errorf("unrecognized op '%s'", op))
			}
			instrs = append(instrs, instr{
				op:  op,
				a:   a,
				b:   b,
				dst: dst,
			})
		}
	}
	return instrs
}
