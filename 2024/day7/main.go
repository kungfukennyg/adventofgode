package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kungfukennyg/adventofgode/common"
)

type op int

const (
	opAdd op = iota
	opMult
	opConcat
)

func (o op) apply(a, b int) int {
	switch o {
	case opAdd:
		return a + b
	case opMult:
		return a * b
	case opConcat:
		return common.MustAtoi(strconv.Itoa(a) + strconv.Itoa(b))
	default:
		panic(fmt.Sprintf("unexpected day7.op: %#v", o))
	}
}

type eq struct {
	want int
	nums []int
}

func (e eq) solve(ops []op) int {
	n := e.nums[0]
	for i, o := range ops {
		n = o.apply(n, e.nums[i+1])
	}

	return n
}

func permute(ops []op, n int) [][]op {
	if n < 2 {
		return nil
	}
	total := 1
	for range n - 1 {
		total *= len(ops)
	}

	perms := [][]op{}
	for i := range total {
		num := i
		line := []op{}
		for range n - 1 {
			opIdx := num % len(ops)
			line = append(line, ops[opIdx])
			num /= len(ops)
		}
		perms = append(perms, line)
	}
	return perms
}

func generatePermutations(ops []op, n int) [][][]op {
	out := [][][]op{}
	for i := 2; i <= n; i++ {
		out = append(out, permute(ops, i))
	}
	return out
}

func (e eq) isPossible(perms [][]op) bool {
	for _, ops := range perms {
		if e.solve(ops) == e.want {
			return true
		}
	}

	return false
}

func calibrationResult(input string, ops []op) int {
	eqs := parse(input)
	longest := 0
	for _, e := range eqs {
		longest = max(len(e.nums), longest)
	}
	perms := generatePermutations(ops, longest)

	possible := []int{}
	for _, e := range eqs {
		if e.isPossible(perms[len(e.nums)-2]) {
			possible = append(possible, e.want)
		}
	}

	sum := 0
	for _, n := range possible {
		sum += n
	}
	return sum
}

func parse(input string) []eq {
	eqs := []eq{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		want := common.MustAtoi(parts[0])
		nums := common.Ints(parts[1], " ")
		eqs = append(eqs, eq{want, nums})
	}
	return eqs
}
