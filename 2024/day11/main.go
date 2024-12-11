package day11

import (
	"maps"
	"strings"

	"github.com/kungfukennyg/adventofgode/aoc"
)

var powersOfTen = []int{
	1, 10, 100, 1000, 10000, 100000, 1000000,
	10000000, 100000000, 1000000000, 10000000000,
}

func powerOfTen(n int) int {
	if n < len(powersOfTen) {
		return powersOfTen[n]
	}

	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func calcDigits(i int) int {
	if i < 10 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func blink(stones map[int]int) map[int]int {
	next := maps.Clone(stones)
	for s, v := range stones {
		next[s] -= v
		if s == 0 {
			next[1] += v
		} else if digits := calcDigits(s); digits%2 == 0 {
			half := digits / 2
			power := powerOfTen(half)

			// Split the number
			a := s / power
			b := s % power
			next[a] += v
			next[b] += v
		} else {
			next[s*2024] += v
		}
	}
	for k, v := range next {
		if v < 1 {
			delete(next, k)
		}
	}
	return next
}

func sim(input string, steps int) int {
	stones := parse(input)
	for range steps {
		stones = blink(stones)
	}
	sum := 0
	for _, s := range stones {
		sum += s
	}
	return sum
}

func parse(input string) map[int]int {
	parts := strings.Split(input, " ")
	stones := map[int]int{}
	for _, s := range parts {
		stones[aoc.MustAtoi(s)] += 1
	}
	return stones
}
