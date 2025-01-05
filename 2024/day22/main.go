package day22

import (
	"math"

	"github.com/kungfukennyg/adventofgode/aoc"
)

const prune int64 = 1 << 24

func next(cur int64) int64 {
	// fast multiply, cur << k = cur * pow(2, k)
	cur ^= (cur << 6)
	// fast modulo
	cur &= prune - 1

	cur ^= (cur >> 5)
	cur &= prune - 1

	cur ^= (cur << 11)
	cur &= prune - 1

	return cur
}

func lastDigit(n int64) int {
	return int(n % 10)
}

type seq [4]int

func validSequences(prices []int) map[seq]int {
	seqs := map[seq]int{}
	for i := 4; i < len(prices); i++ {
		sub := prices[i-3 : i+1]
		seq := seq{}
		for j, p := range sub {
			seq[j] = p - prices[i-4+j]
		}

		if _, ok := seqs[seq]; ok {
			continue
		}

		seqs[seq] = prices[i]
	}
	return seqs
}

func getPrices(secret int64, depth int) []int {
	prices := make([]int, 0, depth+1)
	prices = append(prices, lastDigit(secret))
	num := secret
	for range depth {
		num = next(num)
		ld := lastDigit(num)
		prices = append(prices, ld)
	}

	return prices
}

func mostBananas(input string, depth int) int {
	secrets := aoc.Int64s(input, "\n")

	maxBananas := math.MinInt64
	seqs := map[seq]int{}
	for _, secret := range secrets {
		prices := getPrices(secret, depth)

		sequences := validSequences(prices)

		for seq, cost := range sequences {
			seqs[seq] += cost
			maxBananas = max(maxBananas, seqs[seq])
		}
	}

	return maxBananas
}

func sumSecrets(input string, depth int) int64 {
	secrets := aoc.Int64s(input, "\n")
	for range depth {
		for i, sec := range secrets {
			secrets[i] = next(sec)
		}
	}

	var sum int64
	for _, sec := range secrets {
		sum += sec
	}
	return sum
}
