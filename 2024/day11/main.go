package day11

import (
	"log"
	"strings"
	"sync"

	"github.com/kungfukennyg/adventofgode/common"
)

const concurrency = 2048

func blink(stones []int) []int {
	var wg sync.WaitGroup
	concurrency := min(concurrency, len(stones))
	wg.Add(concurrency)

	bufs := make([][]int, concurrency)
	for i := 0; i < concurrency; i++ {
		j := i

		divide := len(stones) / concurrency
		end := (i + 1) * divide
		if end > len(stones) || (i+1 == concurrency && end < len(stones)) {
			end = len(stones)
		}
		sub := stones[i*divide : end]

		go func(stones []int, i int) {
			out := make([]int, 0, len(stones)*2)
			for _, s := range stones {
				if s == 0 {
					out = append(out, 1)
					continue
				}

				digits := calcDigits(s)
				if digits%2 == 0 {
					half := digits / 2
					power := powerOfTen(half)

					// Split the number
					a := s / power
					b := s % power
					out = append(out, a, b)
					continue
				}

				out = append(out, s*2024)
			}
			bufs[i] = out
			wg.Done()
		}(sub, j)

	}
	wg.Wait()
	out := make([]int, 0, len(stones)*2)
	for _, o := range bufs {
		out = append(out, o...)
	}
	return out
}

// Precomputed powers of 10 for efficiency
var powersOfTen = []int{
	1, 10, 100, 1000, 10000, 100000, 1000000,
	10000000, 100000000, 1000000000,
}

func powerOfTen(n int) int {
	if n < len(powersOfTen) {
		return powersOfTen[n]
	}

	// For very large n, fallback to a loop-based computation
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func calcDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func sim(input string, steps int) int {
	stones := parse(input)
	for i := range steps {
		log.Println(i)
		stones = blink(stones)
	}
	return len(stones)
}

func parse(input string) []int {
	parts := strings.Split(input, " ")
	stones := make([]int, len(parts))
	for i, s := range parts {
		stones[i] = common.MustAtoi(s)
	}
	return stones
}
