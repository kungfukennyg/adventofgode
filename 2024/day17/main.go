package day17

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
	"github.com/kungfukennyg/adventofgode/aoc/opt"
)

var instrNames = []string{"adv", "bxl", "bst", "jnz", "bxc", "out", "bdv", "cdv"}

type instruction uint8

const (
	adv instruction = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func (i instruction) isLiteral() bool {
	switch i {
	case bxl:
		return true
	case bxc:
		return true
	case jnz:
		return true
	default:
		return false
	}
}

func (i instruction) String() string {
	return instrNames[i]
}

type operand uint8

const (
	Literal  operand = 0
	A        operand = 4
	B        operand = 5
	C        operand = 6
	Reserved operand = 7
)

func operandFromValue(v uint8) operand {
	if v <= 3 {
		return Literal
	}

	return operand(v)
}

type cpu struct {
	a, b, c uint64
	pc      uint64
	memory  []uint8
	output  []uint8
}

func (c *cpu) read(n uint8) uint64 {
	op := operandFromValue(n)
	switch op {
	case Literal:
		return uint64(n)
	case A:
		return c.a
	case B:
		return c.b
	case C:
		return c.c
	default:
		panic(fmt.Errorf("cpu.read: unexpected operand '%v'", op))
	}
}

func (c *cpu) Clone() *cpu {
	return &cpu{
		a:      c.a,
		b:      c.b,
		c:      c.c,
		pc:     c.pc,
		memory: slices.Clone(c.memory),
		output: slices.Clone(c.output),
	}
}

func (c *cpu) Run() string {
	for c.step() {

	}

	return aoc.JoinUint8s(c.output, ",")
}

func (c *cpu) step() bool {
	if c.pc+1 >= uint64(len(c.memory)) {
		return false
	}

	instr := instruction(c.memory[c.pc])
	var operand uint64
	if instr.isLiteral() {
		operand = uint64(c.memory[c.pc+1])
	} else {
		operand = c.read(c.memory[c.pc+1])
	}

	switch instr {
	case adv:
		// a / pow(2, operand)
		c.a >>= operand
	case bdv:
		// b = a / pow(2, operand)
		c.b = c.a >> operand
	case cdv:
		// c = a / pow(2, operand)
		c.c = c.a >> operand
	case bxl:
		// bitwise XOR
		c.b ^= operand
	case bst:
		// modulo 8, aka keep lowest 3 bits, 0x7 = 0b0000_0111
		c.b = operand & 0x7
	case jnz:
		// jump if not zero
		if c.a != 0 {
			c.pc = operand
			return true
		}
	case bxc:
		// bitwise XOR b and c
		c.b ^= c.c
	case out:
		// write 3 lowest bits to output
		o := operand & 0x7
		c.output = append(c.output, uint8(o))
	}

	c.pc += 2
	return true
}

func inputFromProgram(input string) uint64 {
	c := parse(input)
	var backtrack func(a uint64, index int) opt.Some[uint64]
	backtrack = func(a uint64, index int) opt.Some[uint64] {
		var bit uint64
		for bit = range 8 {
			c := c.Clone()
			next := (a << 3) | bit
			c.a = next
			c.Run()
			if c.output[0] == c.memory[index] {
				if index == 0 {
					return opt.From(next)
				}

				if v := backtrack(next, index-1); v.Ok() {
					return v
				}
			}
		}
		return opt.None[uint64]()
	}

	return backtrack(0, len(c.memory)-1).Or(0)
}

func parse(input string) *cpu {
	c := cpu{output: []uint8{}}
	for _, line := range aoc.Lines(input) {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Split(line, ": ")
		if strings.HasPrefix(line, "Register") {
			reg := strings.ReplaceAll(parts[0], "Register ", "")
			switch reg {
			case "A":
				c.a = uint64(aoc.MustAtoi(parts[1]))
			case "B":
				c.b = uint64(aoc.MustAtoi(parts[1]))
			case "C":
				c.c = uint64(aoc.MustAtoi(parts[1]))
			}
		} else if strings.HasPrefix(line, "Program:") {
			c.memory = aoc.Uint8s(parts[1], ",")
		}
	}
	return &c
}
